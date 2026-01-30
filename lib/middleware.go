package lib

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"nproxy/config"

	"github.com/go-chi/chi/v5"
)

var mwClient http.Client = http.Client{}

var ROUTE_LOGIN = "/auth/login"
var ROUTE_API = "/api/*path"
var SUBSONIC_REST = "/rest/*path"

// Global and root structure
type NavidromeMiddleware struct {
	Provider *NavidromeExtProvider
	Server   *chi.Mux
}

func NewMiddleware(provider *NavidromeExtProvider) *NavidromeMiddleware {
	server := chi.NewRouter()
	slog.Info(fmt.Sprintf("Using provider: %s", provider.Name))

	// Forward all requests by default then if user
	// registers it's own provider proceed to override
	server.Handle("/*", ForwardMiddleware())

	// Register all handlers from the given provider
	for _, pr := range provider.Handlers {
		slog.Info(fmt.Sprintf("Registering handler for endpoint: %s", pr.SrcPath))
		server.Method(pr.Method, pr.SrcPath, pr.HandlerWithContext())
	}

	return &NavidromeMiddleware{
		Provider: provider,
		Server:   server,
	}
}

// TODO: add options for host and port for example
func (mw *NavidromeMiddleware) RunServer() {
	slog.Info("Listening on port 3000...")
	if err := http.ListenAndServe(":3000", mw.Server); err != nil {
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

func ForwardMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := ForwardRequest(w, r)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
