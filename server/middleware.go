package server

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"nproxy/config"
	"strings"

	"github.com/go-chi/chi/v5"
)

var mwClient http.Client = http.Client{}

// Global and root structure
type NavidromeMiddleware struct {
	Provider *NavidromeExtProvider
	Server   *chi.Mux
}

type NavidromeExtProviderOptions struct {
	// I make available a function to handle the /<PROVIDER>/download route
	// which is created by default when the middleware is created
	// if nil then the route is not created at all
	DownloadHandler func(chan map[string]int) http.HandlerFunc
}

func NewMiddleware(
	provider *NavidromeExtProvider,
	options *NavidromeExtProviderOptions,
) *NavidromeMiddleware {
	server := chi.NewRouter()
	slog.Info(fmt.Sprintf("Using provider: %s", provider.Name))

	// Forward all requests by default then if user
	// registers it's own provider proceed to override
	server.Handle("/*", ForwardMiddleware())

	if options != nil && options.DownloadHandler != nil {
		path := fmt.Sprintf("/%s/download", strings.ToLower(provider.Name))
		slog.Info(fmt.Sprintf("Registering handler for server download at %s", path))
		slog.Info(fmt.Sprintf("Registering handler for server download events at %s", path+"/events"))

		downloadsChan := make(chan map[string]int)

		// If we have a download handler then register it
		server.Handle(path, options.DownloadHandler(downloadsChan))
		// Along with the handler to download we need the one to track the download progres
		// for this we can use SSE (Server Sent Events)
		server.Handle(path+"/events", SseHandler(downloadsChan))
	}

	// Register all handlers from the given provider
	for _, pr := range provider.Handlers {
		slog.Info(fmt.Sprintf("Registering handler for endpoints: %v", pr.SrcPaths))
		for _, srcPath := range pr.SrcPaths {
			server.Method(pr.Method, srcPath, pr.Handler)
		}
	}

	return &NavidromeMiddleware{
		Provider: provider,
		Server:   server,
	}
}

func (mw *NavidromeMiddleware) AddRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) {
	mw.Server.Method(method, path, handler)
}

func (mw *NavidromeMiddleware) RunServer() {
	host := config.GetHost()
	port := config.GetPort()

	slog.Info(fmt.Sprintf("Listening on %s:%s...", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mw.Server); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func (mw *NavidromeMiddleware) UseProvider(provider *NavidromeExtProvider) {
	if provider == nil {
		panic("You must provide a valid NavidromeExtProvider pointer")
	}

	mw.Provider = provider
}

// Forward a request to default navidrome without intercepting it
func ForwardRequest(w http.ResponseWriter, r *http.Request) []byte {
	targetUrl := fmt.Sprintf("%s%s", config.GetNavidromeUrl(), r.URL.Path)
	forwardRequest, err := http.NewRequest(r.Method, targetUrl, r.Body)
	CheckError(w, err)

	for key, values := range r.Header {
		value := values[0]
		forwardRequest.Header.Add(key, value)
	}

	params := r.URL.Query()
	currentQuery := forwardRequest.URL.Query()
	for k, v := range params {
		currentQuery.Add(k, v[0])
	}
	forwardRequest.URL.RawQuery = currentQuery.Encode()

	// TODO: improve error, also sending the right status code
	res, err := mwClient.Do(forwardRequest)
	CheckError(w, err)

	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	body, err := io.ReadAll(res.Body)
	CheckError(w, err)

	w.WriteHeader(res.StatusCode)

	return body
}

// Middleware used to forward requests of endpoint i'm not intercepting (not acting as middlware on them)
func ForwardMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[+] Forwarding request %s\n", r.URL.String())

		body := ForwardRequest(w, r)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
