package handlers

import (
	"encoding/json"
	"net/http"
	"nproxy/lib"
	"nproxy/middlewares/dabmusic/client"
	"nproxy/middlewares/dabmusic/utils"
	"strings"
)

func GetAlbum(
	w http.ResponseWriter,
	r *http.Request,
	client *client.DabClient,
) {
	albumId := r.URL.Query().Get("id")

	// If it's not an external resource than just forward the request and return
	if albumId == "" || !strings.Contains(albumId, "ext-") {
		res := lib.ForwardRequest(w, r)
		w.Write(res)
		return
	}

	splits := strings.Split(albumId, "-")
	parsedAlbumId := splits[2]

	album, err := client.GetAlbumInfo(parsedAlbumId)

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	res := utils.DabToSubsonicAlbum(*album)
	encoded, err := json.Marshal(res)

	if err != nil {
		error := lib.BuildSubsonicError(0, "Unable to decode response from dabmusic")
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encoded)
}
