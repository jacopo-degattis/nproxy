package lib

type TrackAlbumArtist struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Missing bool   `json:"missing"`
}

type TrackParticipants struct {
	AlbumArtist []TrackAlbumArtist `json:"albumArtist"`
	Artist      []TrackAlbumArtist `json:"artist"`
}

type Track struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	Title    string `json:"title"`
	Album    string `json:"album"`
	ArtistId string `json:"artistId"`
	Artist   string `json:"artist"`
	AlbumId  string `json:"albumId"`
	Year     string `json:"year"`
	Size     int    `json:"size"`
	Suffix   string `json:"suffix"`
	Duration int    `json:"duration"`
	Bitrate  int    `json:"birate"`
	Bitdepth int    `json:"bitdepth"`
}

type NavidromeTrack struct {
	Track

	Bookmarkposition     int               `json:"bookmarkPosition"`
	Libraryid            int               `json:"libraryId"`
	LibraryPath          string            `json:"libraryPath"`
	Libraryname          string            `json:"libraryName"`
	Folderid             string            `json:"folderId"`
	Albumartistid        string            `json:"albumArtistId"`
	Albumartist          string            `json:"albumArtist"`
	HasCoverArt          bool              `json:"hasCoverArt"`
	Tracknumber          int               `json:"trackNumber"`
	Discnumber           int               `json:"discNumber"`
	Date                 int               `json:"date"`
	Samplerate           int               `json:"sampleRate"`
	Channels             int               `json:"channels"`
	Genre                string            `json:"genre"`
	Ordertitle           string            `json:"orderTitle"`
	Orderalbumname       string            `json:"orderAlbumName"`
	Orderartistname      string            `json:"orderArtistName"`
	Orderalbumartistname string            `json:"orderAlbumArtistName"`
	Compilation          bool              `json:"compilation"`
	Lyrics               string            `json:"lyrics"`
	Explicitstatus       string            `json:"explicitStatus"`
	Rgalbumgain          *int              `json:"rgAlbumGain"`
	Rgalbumpeak          *int              `json:"rgAlbumPeak"`
	Rgtrackgain          *int              `json:"rgTrackGain"`
	Rgtrackpeak          *int              `json:"rgTrackPeak"`
	Participants         TrackParticipants `json:"particiants"`
	Missing              bool              `json:"missing"`
	Birthtime            string            `json:"birthTime"`
	Createdat            string            `json:"createdAt"`
	Updatedat            string            `json:"updatedAt"`
}
