package dabmusic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"nproxy/config"
	"nproxy/lib"
	libTypes "nproxy/lib"
	"nproxy/redisdb"
	"strings"

	"github.com/redis/go-redis/v9"
)

var dabClient = DabClient{
	BaseUrl: "https://dabmusic.xyz",
}

func MiddlewareApiLoginNavidrome(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiLoginNavidromeHandler := lib.ProviderHandler{
		Method:   "POST",
		SrcPath:  "/auth/login",
		DstPath:  "", // TODO: this should be also settable to null to avoid automatic forward to that endpoint
		Provider: provider,
	}

	apiLoginNavidromeHandler.Handler = func(w http.ResponseWriter, r *http.Request) {
		fullUrl := fmt.Sprintf("%s%s", config.GetNavidromeUrl(), "/auth/login")

		defer r.Body.Close()
		res, err := lib.Fetch(fullUrl, "POST", r.Body, r.Header)

		if err != nil {
			error := lib.BuildSubsonicError(0, err.Error())
			w.WriteHeader(500)
			w.Write(error)
			return
		}

		rawBody := make(map[string]any)
		err = json.NewDecoder(bytes.NewReader(res)).Decode(&rawBody)

		if err != nil {
			error := lib.BuildSubsonicError(0, err.Error())
			w.WriteHeader(500)
			w.Write(error)
			return
		}

		var token = rawBody["token"].(string)
		ctx := context.WithValue(r.Context(), "token", token)

		r = r.WithContext(ctx)

		w.Write(res)
	}

	return apiLoginNavidromeHandler
}

// Naming format: Middleware<ROUTE_PREFIX><ROUTE_POSTFIX><SERVICE>
// Example: /api/song = MiddlewareApiSongNavidrome
func MiddlewareApiSongNavidrome(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	type SearchResponse struct {
		Tracks []DabTrack `json:"tracks"`
	}

	apiSongNavidromeHandler := lib.ProviderHandler{
		Method:   "GET",
		SrcPath:  "/api/song",
		DstPath:  "/api/search",
		Provider: provider,
	}

	apiSongNavidromeHandler.Handler = func(w http.ResponseWriter, r *http.Request) {
		// var dstPath = r.Context().Value("DstPath").(string)

		userQuery := r.URL.Query().Get("title")

		if userQuery == "" {
			// If the client doesn't provide any title to query than just return
			// the default navidrome response
			res := lib.ForwardRequest(w, r)
			w.Write(res)
			return
		}

		tracks, err := dabClient.Search(
			userQuery,
			"track",
		)

		if err != nil {
			error := lib.BuildSubsonicError(0, err.Error())
			w.WriteHeader(500)
			w.Write(error)
			return
		}

		if len(tracks) == 0 {
			res := lib.ForwardRequest(w, r)
			w.Write(res)
			return
		}

		// convert tracks to navidrome compatible format
		navidromeTracks := make([]libTypes.NavidromeTrack, 0)
		for _, dabTrack := range tracks {
			navidromeTracks = append(navidromeTracks, DabToNavidromeTrack(dabTrack))
		}
		encoded, err := json.Marshal(navidromeTracks)

		if err != nil {
			error := lib.BuildSubsonicError(0, "Unable to decode response from dabmusic")
			w.WriteHeader(500)
			w.Write(error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(encoded)
	}

	return apiSongNavidromeHandler
}

// TODO, NOTE: If I have two endpoints like `/rest/getCoverArt` and `/rest/getCovertArtView` which are similar or
// even equal i can create two different providers and assign the same handler to them
func MiddlewareRestGetCoverArtView(provider *lib.NavidromeExtProvider) lib.ProviderHandler {
	apiGetCoverArtSubsonic := lib.ProviderHandler{
		Method:   "GET",
		SrcPath:  "/rest/getCoverArt.view",
		DstPath:  "", // TODO: this should be also settable to null to avoid automatic forward to that endpoint
		Provider: provider,
	}

	apiGetCoverArtSubsonic.Handler = func(w http.ResponseWriter, r *http.Request) {
		// Id format for external tracks is: ext-<PROVIDER>-<EXT_TRACK_ID>

		queryId := r.URL.Query().Get("id")

		isExternal := strings.Contains(queryId, "ext")

		// If not external id then let navidrome handle it
		if !isExternal {
			res := lib.ForwardRequest(w, r)
			w.Write(res)
			return
		}

		// otherwise we handle it now
		externalTrackId := strings.Split(queryId, "-")

		redisCmd := redisdb.Get(externalTrackId[2])

		if !errors.Is(redisCmd.Err(), redis.Nil) {
			cachedImg, err := redisCmd.Bytes()
			if err != nil {
				error := lib.BuildSubsonicError(0, err.Error())
				w.Write(error)
				return
			}
			w.Header().Set("Content-Type", "image/png")
			w.Write(cachedImg)
			return
		}

		metadata, err := dabClient.GetTrackMetadata(externalTrackId[2])

		if err != nil {
			error := lib.BuildSubsonicError(0, err.Error())
			w.Write(error)
			return
		}

		metadataBytes, err := lib.Fetch(metadata.Cover, "GET", nil, nil)

		// Store the img in cache now
		redisdb.Set(externalTrackId[2], metadataBytes)

		if err != nil {
			error := lib.BuildSubsonicError(0, err.Error())
			w.Write(error)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(metadataBytes)
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
	apiGetCoverArtSubsonicHandler := MiddlewareRestGetCoverArtView(&dabMusicProvider)

	dabMusicProvider.Handlers = []lib.ProviderHandler{
		apiLoginNavidromeHandler,
		apiSongNavidromeHandler,
		apiGetCoverArtSubsonicHandler,
	}

	return &dabMusicProvider
}
