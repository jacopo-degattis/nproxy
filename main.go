package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"nproxy/config"
	"nproxy/redisdb"
	lib "nproxy/server"
	"os"

	downloader "nproxy/downloader"
	dabmusicMiddleware "nproxy/middlewares/dabmusic"
	"nproxy/middlewares/dabmusic/handlers"
)

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
	mw := lib.NewMiddleware(dabMusicProvider, &lib.NavidromeExtProviderOptions{
		DownloadHandler: func(w http.ResponseWriter, r *http.Request) {
			handlers.DownloadHandler(w, r, dw)
		},
	})
	mw.RunServer()
}
