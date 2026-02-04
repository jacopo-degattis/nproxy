package handlers

import (
	"encoding/json"
	"net/http"
	"nproxy/middlewares/dabmusic/client"
	utils "nproxy/middlewares/dabmusic/utils"
	"nproxy/server"
	libTypes "nproxy/server"
)

// api/song

func ApiSong(
	w http.ResponseWriter,
	r *http.Request,
	client *client.DabClient,
) {
	userQuery := r.URL.Query().Get("title")

	if userQuery == "" {
		// If the client doesn't provide any title to query than just return
		// the default navidrome response
		res := server.ForwardRequest(w, r)
		w.Write(res)
		return
	}

	tracks, err := client.Search(
		userQuery,
		"track",
	)

	if err != nil {
		error := server.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	if len(tracks) == 0 {
		res := server.ForwardRequest(w, r)
		w.Write(res)
		return
	}

	// convert tracks to navidrome compatible format
	responseTracks := make([]libTypes.NavidromeTrack, 0)
	for _, dabTrack := range tracks {
		responseTracks = append(responseTracks, utils.DabToNavidromeTrack(dabTrack))
	}
	encoded, err := json.Marshal(responseTracks)

	if err != nil {
		error := server.BuildSubsonicError(0, "Unable to decode response from dabmusic")
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encoded)
}
