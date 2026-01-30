package dabmusic

type DabTrack struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"albumTitle"`
	Cover       string `json:"albumCover"`
	ReleaseDate string `json:"releaseDate"`
	Duration    int    `json:"duration"`
	TrackNumber int    `json:"-"`
}
