package middleware

import (
	"fmt"
	"io"
	"log"
	"navidrome-middleware/config"
	"net/http"

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

	// Forward all requests by default then if user
	// registers it's own provider proceed to override
	server.Handle("/*", ForwardMiddleware())

	// Register all handlers from the given provider
	for _, pr := range provider.Handlers {
		switch pr.Method {
		case "GET":
			server.Get(pr.SrcPath, pr.HandlerWithContext())
		case "POST":
			server.Post(pr.SrcPath, pr.HandlerWithContext())
		}
	}

	return &NavidromeMiddleware{
		Provider: provider,
		Server:   server,
	}
}

// TODO: add options for host and port for example
func (mw *NavidromeMiddleware) RunServer() {
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

func ForwardMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetUrl := fmt.Sprintf("%s%s", config.GetNavidromeUrl(), r.URL)

		forwardRequest, err := http.NewRequest(r.Method, targetUrl, r.Body)
		CheckError(w, err)
		res, err := mwClient.Do(forwardRequest)
		CheckError(w, err)

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		CheckError(w, err)

		w.Write(body)
	}
}
