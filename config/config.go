package config

import (
	"os"
)

var Env = map[string]string{
	// Host is the host to run the server onto (ex. localhost)
	"HOST": os.Getenv("HOST"),
	// Port is the port to run the server onto (ex. 3000)
	"PORT":           os.Getenv("PORT"),
	"REDIS_URL":      os.Getenv("REDIS_URL"),
	"REDIS_PASSWORD": os.Getenv("REDIS_PASSWORD"),
	"NAVIDROME_URL":  os.Getenv("NAVIDROME_URL"),
	"DAB_ENDPOINT":   os.Getenv("DAB_ENDPOINT"),
	"VERSION":        "1.0.0",
}

func GetHost() string {
	host := Env["HOST"]
	if host != "" {
		return host
	}
	return "localhost"
}

func GetPort() string {
	port := Env["PORT"]
	if port != "" {
		return port
	}
	return "3000"
}

func GetRedisUrl() string {
	redisUrl := Env["REDIS_URL"]
	if redisUrl != "" {
		return redisUrl
	}
	return "localhost:6379"
}

func GetRedisPassword() *string {
	redisPassword := Env["REDIS_PASSWORD"]
	if redisPassword != "" {
		return &redisPassword
	}
	return nil
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
