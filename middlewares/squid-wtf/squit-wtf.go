package squidwtf

import (
	"net/http"
	handlers "nproxy/middlewares/squid-wtf/handlers"
	server "nproxy/server"
)

func SquidWtfProvider() *server.NavidromeExtProvider {
	squidWtfProvider := server.NavidromeExtProvider{
		Name:    "squid-wtf",
		BaseUrl: "https://qobuz.squid.wtf",
	}

	apiSongNavidromeHandler := squidWtfProvider.CreateProviderHandler(
		"GET",
		[]string{"/api/song"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.ApiSong(w, r)
		},
	)

	squidWtfProvider.Handlers = []server.ProviderHandler{
		apiSongNavidromeHandler,
	}

	return &squidWtfProvider
}
