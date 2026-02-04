package types

import (
	"strconv"
	"strings"
)

type ID int

func (id *ID) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, `"`)
	val, _ := strconv.Atoi(s)
	*id = ID(val)
	return nil
}

type DabAlbum struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Artist      string     `json:"artist"`
	ReleaseDate string     `json:"releaseDate"`
	Genre       string     `json:"genre"`
	Cover       string     `json:"cover"`
	Tracks      []DabTrack `json:"tracks"`
}

type DabTrack struct {
	Id          ID     `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ArtistId    int    `json:"artistId"`
	Album       string `json:"albumTitle"`
	AlbumId     string `json:"albumId"`
	Cover       string `json:"albumCover"`
	ReleaseDate string `json:"releaseDate"`
	Duration    int    `json:"duration"`
	TrackNumber int    `json:"-"`
}
