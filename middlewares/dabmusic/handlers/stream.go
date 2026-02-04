package handlers

import (
	"fmt"
	"io"
	"net/http"
	"nproxy/lib"
	"nproxy/middlewares/dabmusic/client"
	"strings"
)

func Stream(
	w http.ResponseWriter,
	r *http.Request,
	client *client.DabClient,
) {
	fmt.Println("/rest/stream requested")
	fmt.Printf("Plain url: %s\n", r.URL.String())

	trackId := r.URL.Query().Get("id")

	isExternal := strings.Contains(trackId, "ext")

	if !isExternal {
		res := lib.ForwardRequest(w, r)
		w.Write(res)
		return
	}

	externalTrackId := strings.Split(trackId, "-")

	streamUrl, err := client.GetTrackStreamUrl(externalTrackId[2])

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	// res, err := lib.Fetch(*streamUrl, "GET", nil, r.Header, nil)
	cl := http.Client{}
	req, err := http.NewRequest("GET", *streamUrl, nil)

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	streamResponse, err := cl.Do(req)

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	defer streamResponse.Body.Close()
	resBody, err := io.ReadAll(streamResponse.Body)

	if err != nil {
		error := lib.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	w.Header().Set("Accept-ranges", streamResponse.Header.Get("Accept-ranges"))
	w.Header().Set("Content-Length", streamResponse.Header.Get("Content-Length"))
	w.Header().Set("Content-Type", streamResponse.Header.Get("Content-Type"))
	w.Write(resBody)
}
