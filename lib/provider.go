package lib

import (
	"context"
	"net/http"
)

type ProviderHandler struct {
	Method   string
	SrcPath  string // path we want to intercept of navidrome or subsonic api
	DstPath  string // path of the route we want to forward our request to
	Handler  http.HandlerFunc
	Provider *NavidromeExtProvider
}

func (ph *ProviderHandler) HandlerWithContext() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add destination path to the context
		// TODO: howto: c.Set("DstPath", ph.DstPath)?
		ctx := context.WithValue(r.Context(), "DstPath", ph.DstPath)
		ph.Handler.ServeHTTP(w, r.WithContext(ctx))
	}
}

type NavidromeExtProvider struct {
	Name     string
	BaseUrl  string
	Handlers []ProviderHandler
}
