package handlers

import (
	"encoding/json"
	"net/http"

	client "nproxy/middlewares/dabmusic/client"
	utils "nproxy/middlewares/dabmusic/utils"
	"nproxy/server"
)

// /rest/search3 & /rest/search3.view
func Search3Handler(
	w http.ResponseWriter,
	r *http.Request,
	client *client.DabClient,
) {

	userQuery := r.URL.Query().Get("query")

	if userQuery == "\"\"" {
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

	res := utils.DabToSubsonicTrack(tracks)
	encoded, err := json.Marshal(res)

	if err != nil {
		error := server.BuildSubsonicError(0, "Unable to decode response from dabmusic")
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encoded)
}
