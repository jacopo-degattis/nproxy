package dabmusic

import (
	"net/http"
	"nproxy/lib"

	client "nproxy/middlewares/dabmusic/client"
	handlers "nproxy/middlewares/dabmusic/handlers"
)

var dabClient = client.DabClient{
	BaseUrl: "https://dabmusic.xyz",
}

func DabMusicProvider() *lib.NavidromeExtProvider {
	dabMusicProvider := lib.NavidromeExtProvider{
		Name:    "dabmusic",
		BaseUrl: "https://dabmusic.xyz",
	}

	apiLoginNavidromeHandler := dabMusicProvider.CreateProviderHandler(
		"POST",
		[]string{"/auth/login"},
		handlers.LoginHandler,
	)

	apiSongNavidromeHandler := dabMusicProvider.CreateProviderHandler(
		"GET",
		[]string{"/api/song"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.ApiSong(w, r, &dabClient)
		},
	)

	apiRestSearch3Subsonic := dabMusicProvider.CreateProviderHandler(
		"GET",
		[]string{"/rest/search3.view", "/rest/search3"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Search3Handler(w, r, &dabClient)
		},
	)

	apiGetCoverArtViewSubsonicHandler := dabMusicProvider.CreateProviderHandler(
		"GET",
		[]string{"/rest/getCoverArt.view", "/rest/getCoverArt"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetCoverArt(w, r, &dabClient)
		},
	)

	apiStreamSubsonicHandler := dabMusicProvider.CreateProviderHandler(
		"GET",
		[]string{"/rest/stream"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Stream(w, r, &dabClient)
		},
	)

	apiRestAlbumHandler := dabMusicProvider.CreateProviderHandler(
		"GET",
		[]string{"/rest/getAlbum"},
		func(w http.ResponseWriter, r *http.Request) {
			handlers.GetAlbum(w, r, &dabClient)
		},
	)

	dabMusicProvider.Handlers = []lib.ProviderHandler{
		apiLoginNavidromeHandler,
		apiSongNavidromeHandler,
		apiRestSearch3Subsonic,
		apiGetCoverArtViewSubsonicHandler,
		apiStreamSubsonicHandler,
		apiRestAlbumHandler,
	}

	return &dabMusicProvider
}
