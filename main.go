package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"nproxy/config"
	"nproxy/lib"
	"nproxy/redisdb"
	"os"

	downloader "nproxy/downloader"
	dabmusicMiddleware "nproxy/middlewares/dabmusic"
	dmClient "nproxy/middlewares/dabmusic/client"

	"go.senan.xyz/taglib"
)

type Metadatas struct {
	Title       string
	Artist      string
	Album       string
	Date        string
	Cover       string
	TrackNumber int
}

func addMetadata(targetFile string, metadatas Metadatas) error {
	// res, err := _request(metadatas.Cover, false, []QueryParams{})
	req, err := http.NewRequest("GET", metadatas.Cover, nil)
	if err != nil {
		return fmt.Errorf("can't download cover")
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("can't download cover")
	}
	defer res.Body.Close()

	coverBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = taglib.WriteTags(targetFile, map[string][]string{
		taglib.Title:  {metadatas.Title},
		taglib.Artist: {metadatas.Artist},
		taglib.Album:  {metadatas.Album},
		taglib.Date:   {metadatas.Date},
	}, 0)

	if err != nil {
		return fmt.Errorf("unable to write metadata to track")
	}

	if len(coverBytes) > 0 {
		err = taglib.WriteImage(targetFile, coverBytes)
	}

	if err != nil {
		return fmt.Errorf("unable to picture meadata to track")
	}

	return nil
}

func main() {
	if config.GetNavidromeUrl() == "" {
		log.Fatal("you should set a valid `NAVIDROME_URL` env variable.")
	}

	// Initialize redis db connection
	if err := redisdb.InitRedisClient(); err != nil {
		slog.Error(fmt.Sprintf("unable to connect to redis client on address %s", config.GetRedisUrl()))
		os.Exit(1)
	}

	client := http.Client{}
	dw := downloader.Downloader{
		Client:              &client,
		DownloadDir:         "./",
		MaximumDownloadPool: 1,
	}

	// Define the provider
	dabMusicProvider := dabmusicMiddleware.DabMusicProvider()

	// Now, for example purposes using the dabmusic provider
	mw := lib.NewMiddleware(dabMusicProvider)
	mw.AddRoute("GET", "/dabmusic/download", func(w http.ResponseWriter, r *http.Request) {
		// Route to download tracks or albums (in future also artists) from dabmusic onto the server
		// TODO: as a feature it would be nice to be able to keep track of download progress so find a way

		resId := r.URL.Query().Get("resId")
		downloadType := r.URL.Query().Get("downloadType")

		if resId == "" || downloadType == "" {
			error := "you should provide a valid `downloadType` and `resId` query params"
			w.Write([]byte(error))
			return
		}

		switch downloadType {
		case "album":
			fmt.Print("Handling track")
		case "track":
			client := dmClient.DabClient{
				BaseUrl: "https://dabmusic.xyz",
			}
			trackMetadata, error := client.GetTrackMetadata(resId)

			if error != nil {
				w.Write([]byte(error.Error()))
				return
			}

			streamUrl, error := client.GetTrackStreamUrl(resId)

			if error != nil {
				w.Write([]byte(error.Error()))
				return
			}

			fmt.Println("Starting download now")
			dw.DownloadFrom(*streamUrl, fmt.Sprintf("%s.mp3", trackMetadata.Title))
			addMetadata(fmt.Sprintf("%s.mp3", trackMetadata.Title), Metadatas{
				Title:       trackMetadata.Title,
				Artist:      trackMetadata.Title,
				Album:       trackMetadata.Album,
				Date:        trackMetadata.ReleaseDate,
				Cover:       trackMetadata.Cover,
				TrackNumber: trackMetadata.TrackNumber,
			})
		default:
			error := "downloadType must be either `album` or `track`"
			w.Write([]byte(error))
			return
		}

		w.Write([]byte("DONE"))
	})
	mw.RunServer()

	// client := http.Client{}
	// dw := downloader.Downloader{
	// 	Client:              &client,
	// 	DownloadDir:         "./",
	// 	MaximumDownloadPool: 1,
	// }

	// err := dw.DownloadFrom("https://streaming-qobuz-std.akamaized.net/file?uid=9852090&eid=54114173&fmt=7&profile=raw&app_id=798273057&cid=3747494&etsp=1770230844&hmac=qv4DsQppKUTqYx8Pas3EboKzmjU", "test.mp3")

	// if err != nil {
	// 	panic(err)
	// }
}
