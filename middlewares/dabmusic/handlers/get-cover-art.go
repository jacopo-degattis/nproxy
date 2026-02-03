package handlers

import (
	"errors"
	"net/http"
	"nproxy/lib"
	"nproxy/middlewares/dabmusic/client"
	"nproxy/redisdb"
	"strings"

	"github.com/redis/go-redis/v9"
)

// /rest/getCoverArt & /rest/getCoverArt.view

func GetCoverArt(
	w http.ResponseWriter,
	r *http.Request,
	client *client.DabClient,
) {
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

	metadata, err := client.GetTrackMetadata(externalTrackId[2])

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.Write(error)
		return
	}

	metadataBytes, err := lib.Fetch(metadata.Cover, "GET", nil, nil, nil)

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
