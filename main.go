package main

import (
	"log"
	"nproxy/config"
	"nproxy/lib"
	"nproxy/redisdb"

	dabmusicMiddleware "nproxy/middlewares/dabmusic"
)

// This must stay FIXED in order to support navidrome and subsonic clients
var ROUTE_LOGIN = "/auth/login"

func main() {
	if config.GetNavidromeUrl() == "" {
		log.Fatal("you should set a valid `NAVIDROME_URL` env variable.")
	}

	// Initialize redis db connection
	redisdb.InitRedisClient()

	// Define the provider
	dabMusicProvider := dabmusicMiddleware.DabMusicProvider()

	// Now, for example purposes using the dabmusic provider
	mw := lib.NewMiddleware(dabMusicProvider)
	mw.RunServer()
}
