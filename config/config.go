package config

import (
	"os"
)

var Env = map[string]string{
	"NAVIDROME_URL": os.Getenv("NAVIDROME_URL"),
	"DAB_ENDPOINT":  os.Getenv("DAB_ENDPOINT"),
	"VERSION":       "1.0.0",
}

func GetDabEndpoint() string {
	endpoint := Env["DAB_ENDPOINT"]
	if endpoint != "" {
		return endpoint
	}
	return "https://dabmusic.xyz"
}

func GetNavidromeUrl() string {
	navidromeUrl := Env["NAVIDROME_URL"]
	if navidromeUrl != "" {
		return navidromeUrl
	}
	return ""
}

func GetVersion() string {
	return Env["VERSION"]
}
