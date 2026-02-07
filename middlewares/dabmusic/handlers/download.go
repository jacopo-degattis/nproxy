package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"nproxy/downloader"
	client "nproxy/middlewares/dabmusic/client"
	"nproxy/middlewares/dabmusic/types"
	utils "nproxy/middlewares/dabmusic/utils"
)

func downloadTrack(
	path string,
	trackMetadata *types.DabTrack,
	client *client.DabClient,
	dw *downloader.Downloader,
) error {
	_, err := os.Stat(path + trackMetadata.Title + ".mp3")
	if err == nil {
		return fmt.Errorf("song %s already exists", trackMetadata.Title)
	}

	streamUrl, error := client.GetTrackStreamUrl(strconv.Itoa(int(trackMetadata.Id)))

	if error != nil {
		return error
	}

	fileName := fmt.Sprintf("%s/%s.mp3", path, trackMetadata.Title)
	dw.DownloadFrom(*streamUrl, fileName)
	utils.AddMetadata(fileName, trackMetadata)

	return nil
}

func sanitizeName(name string) string {
	noSpace := strings.TrimSpace(name)
	newName := strings.ToLower(strings.ReplaceAll(noSpace, " ", "_"))
	return newName
}

// Route to download tracks or albums (in future also artists) from dabmusic onto the server
func DownloadHandler(
	w http.ResponseWriter,
	r *http.Request,
	dw downloader.Downloader,
) {
	// TODO: as a feature it would be nice to be able to keep track of download progress so find a way

	resId := r.URL.Query().Get("resId")
	downloadType := r.URL.Query().Get("downloadType")

	if resId == "" || downloadType == "" {
		w.WriteHeader(400)
		error := "you should provide a valid `downloadType` and `resId` query params"
		w.Write([]byte(error))
		return
	}

	client := client.DabClient{
		BaseUrl: "https://dabmusic.xyz",
	}
	switch downloadType {
	case "album":
		fmt.Println("Getting inside the GetAlbumInfo...")
		albumInfos, error := client.GetAlbumInfo(resId)

		fmt.Printf("Downloading: %s\n", albumInfos.Title)

		artistFolder := strings.ToLower(strings.ReplaceAll(albumInfos.Artist, " ", "_"))
		albumFolder := sanitizeName(albumInfos.Title)
		fullAlbumDownloadPath := fmt.Sprintf("%s/%s/", artistFolder, albumFolder)

		if !utils.DirExists(fullAlbumDownloadPath) {
			// Create the path only if it doesn't exists
			err := os.MkdirAll(fullAlbumDownloadPath, os.ModePerm)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(error.Error()))
				return
			}
		}

		// store ids of tracks you can't download
		errors := make([]int, 0)

		for _, t := range albumInfos.Tracks {
			fmt.Printf("Downloading track: %s\n", t.Title)
			error = downloadTrack(fullAlbumDownloadPath, &t, &client, &dw)

			if error != nil {
				errors = append(errors, int(t.Id))
			}
		}

		if len(errors) > 0 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("errors while downloading the following tracks: %v", errors)))
			return
		}
	case "track":
		trackMetadata, error := client.GetTrackMetadata(resId)

		if error != nil {
			w.WriteHeader(500)
			w.Write([]byte(error.Error()))
			return
		}

		artistFolder := strings.ToLower(strings.ReplaceAll(trackMetadata.Artist, " ", "_"))
		albumFolder := strings.ToLower(strings.ReplaceAll(trackMetadata.Title, " ", "_"))
		fullTrackDownloadPath := fmt.Sprintf("%s/%s/", artistFolder, albumFolder)

		error = os.MkdirAll(fullTrackDownloadPath, os.ModePerm)

		if error != nil {
			w.WriteHeader(500)
			w.Write([]byte(error.Error()))
			return
		}

		err := downloadTrack(fullTrackDownloadPath, trackMetadata, &client, &dw)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	default:
		error := "downloadType must be either `album` or `track`"
		w.Write([]byte(error))
		return
	}

	w.WriteHeader(201)
}
