package server

import (
	"net/http"
)

type ProviderHandler struct {
	Method   string
	SrcPaths []string // paths we want to intercept of navidrome or subsonic api
	Handler  http.HandlerFunc
	Provider *NavidromeExtProvider
}

type NavidromeExtProvider struct {
	Name     string
	BaseUrl  string
	Handlers []ProviderHandler
}

func (nxt *NavidromeExtProvider) CreateProviderHandler(
	method string,
	srcPaths []string,
	handlerFunc HttpHandlerFunc,
) ProviderHandler {
	handler := ProviderHandler{
		Method:   method,
		SrcPaths: srcPaths,
		Provider: nxt,
	}

	handler.Handler = handlerFunc

	return handler
}
