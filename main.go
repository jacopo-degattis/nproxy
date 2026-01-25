package main

import (
	"fmt"
	"io"
	"log"
	"navidrome-middleware/config"
	"navidrome-middleware/middleware"
	"net/http"
	"net/url"
)

// This must stay FIXED in order to support navidrome and subsonic clients
var ROUTE_LOGIN = "/auth/login"

func fetch(url string, method string, body io.Reader) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	fmt.Println(res)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return resBody, nil
}

func main() {
	if config.GetNavidromeUrl() == "" {
		log.Fatal("you should set a valid `NAVIDROME_URL` env variable.")
	}

	dabMusicProvider := middleware.NavidromeExtProvider{
		Name:    "dabmusic",
		BaseUrl: "https://dabmusic.xyz",
	}

	loginHandler := middleware.ProviderHandler{
		Method:   "POST",
		SrcPath:  "/api/song",
		DstPath:  "/api/search",
		Provider: &dabMusicProvider,
	}

	loginHandler.Handler = func(w http.ResponseWriter, r *http.Request) {
		var dstPath = r.Context().Value("DstPath").(string)

		userQuery := r.URL.Query().Get("title")
		encodedQuery := url.QueryEscape(userQuery)
		fullPath := fmt.Sprintf("%s%s?q=%s", loginHandler.Provider.BaseUrl, dstPath, encodedQuery)

		fmt.Printf("FULL PATh = %s\n", fullPath)

		res, err := fetch(fullPath, "GET", nil)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(res))

		w.Write(res)
	}
	dabMusicProvider.Handlers = []middleware.ProviderHandler{loginHandler}

	// Define the provider

	// Now, for example purposes using the dabmusic provider
	mw := middleware.NewMiddleware(&dabMusicProvider)
	mw.RunServer()
}
