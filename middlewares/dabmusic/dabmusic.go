package dabmusic

import (
	"net/http"
	"nproxy/lib"

	client "nproxy/middlewares/dabmusic/client"
	handlers "nproxy/middlewares/dabmusic/handlers"
	dabtypes "nproxy/middlewares/dabmusic/types"
)

var dabClient = client.DabClient{
	BaseUrl: "https://dabmusic.xyz",
}

func MiddlewareApiLoginNavidrome(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiLoginNavidromeHandler := lib.ProviderHandler{
		Method:   "POST",
		SrcPaths: []string{"/auth/login"},
		DstPath:  "", // TODO: this should be also settable to null to avoid automatic forward to that endpoint
		Provider: provider,
	}

	apiLoginNavidromeHandler.Handler = handlers.LoginHandler

	return apiLoginNavidromeHandler
}

func MiddlewareRestSearch3Subsonic(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiRestSearch3Subsonic := lib.ProviderHandler{
		Method:   "GET",
		SrcPaths: []string{"/rest/search3.view", "/rest/search3"},
		DstPath:  "/api/search",
		Provider: provider,
	}

	apiRestSearch3Subsonic.Handler = func(w http.ResponseWriter, r *http.Request) {
		handlers.Search3Handler(w, r, &dabClient)
	}

	return apiRestSearch3Subsonic
}

// Naming format: Middleware<ROUTE_PREFIX><ROUTE_POSTFIX><SERVICE>
// Example: /api/song = MiddlewareApiSongNavidrome
func MiddlewareApiSongNavidrome(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	type SearchResponse struct {
		Tracks []dabtypes.DabTrack `json:"tracks"`
	}

	apiSongNavidromeHandler := lib.ProviderHandler{
		Method:   "GET",
		SrcPaths: []string{"/api/song"},
		DstPath:  "/api/search",
		Provider: provider,
	}

	apiSongNavidromeHandler.Handler = func(w http.ResponseWriter, r *http.Request) {
		handlers.ApiSong(w, r, &dabClient)
	}

	return apiSongNavidromeHandler
}

func MiddlwareRestStreamSubsonic(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiRestStreamSubsonic := lib.ProviderHandler{
		Method:   "GET",
		SrcPaths: []string{"/rest/stream"},
		DstPath:  "/api/stream",
		Provider: provider,
	}

	apiRestStreamSubsonic.Handler = func(w http.ResponseWriter, r *http.Request) {
		handlers.Stream(w, r, &dabClient)
	}

	return apiRestStreamSubsonic
}

// TODO, NOTE: If I have two endpoints like `/rest/getCoverArt` and `/rest/getCovertArtView` which are similar or
// even equal i can create two different providers and assign the same handler to them
func MiddlewareRestGetCoverArtView(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiGetCoverArtSubsonic := lib.ProviderHandler{
		Method:   "GET",
		SrcPaths: []string{"/rest/getCoverArt.view", "/rest/getCoverArt"},
		DstPath:  "", // TODO: this should be also settable to null to avoid automatic forward to that endpoint
		Provider: provider,
	}

	apiGetCoverArtSubsonic.Handler = func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCoverArt(w, r, &dabClient)
	}

	return apiGetCoverArtSubsonic
}

func DabMusicProvider() *lib.NavidromeExtProvider {
	dabMusicProvider := lib.NavidromeExtProvider{
		Name:    "dabmusic",
		BaseUrl: "https://dabmusic.xyz",
	}

	apiLoginNavidromeHandler := MiddlewareApiLoginNavidrome(&dabMusicProvider)
	apiSongNavidromeHandler := MiddlewareApiSongNavidrome(&dabMusicProvider)
	apiRestSearch3Subsonic := MiddlewareRestSearch3Subsonic(&dabMusicProvider)
	apiGetCoverArtViewSubsonicHandler := MiddlewareRestGetCoverArtView(&dabMusicProvider)
	apiStreamSubsonicHandler := MiddlwareRestStreamSubsonic(&dabMusicProvider)

	dabMusicProvider.Handlers = []lib.ProviderHandler{
		apiLoginNavidromeHandler,
		apiSongNavidromeHandler,
		apiRestSearch3Subsonic,
		apiGetCoverArtViewSubsonicHandler,
		apiStreamSubsonicHandler,
	}

	return &dabMusicProvider
}
