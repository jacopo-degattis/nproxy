package server

import "net/http"

// Type for http.HandlerFunc
type HttpHandlerFunc = func(w http.ResponseWriter, r *http.Request)

type SubsonicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SubsonicResponse struct {
	Status        string `json:"status"`
	Version       string `json:"version"`
	Type          string `json:"type"`
	ServerVersion string `json:"serverVersion"`
	OpenSubsonic  bool   `json:"openSubsonic"`
}

type SubsonicResponseAlbumDto struct {
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	Artist              string   `json:"artist"`
	ArtistId            string   `json:"artistId"`
	CoverArt            string   `json:"coverArt"`
	SongCount           int      `json:"songCount"`
	Duration            int      `json:"duration"`
	PlayCount           int      `json:"playCount"`
	Created             string   `json:"created"`
	Year                int      `json:"year"`
	Played              string   `json:"played"`
	UserRating          int      `json:"userRating"`
	Genres              []string `json:"genres"`
	MusicBrainzId       string   `json:"musicBrainzId"`
	IsCompilation       bool     `json:"isCompilation"`
	SortName            string   `json:"sortName"`
	DiscTitles          []string `json:"discTitles"`
	OriginalReleaseDate any      `json:"originalReleaseDate"`
	ReleaseDate         any      `json:"releaseDate"`
	ReleaseTypes        []any    `json:"releaseTypes"`
	RecordLabels        []string `json:"recordLabels"`
	Moods               []string `json:"moods"`
	Artists             []any    `json:"artists"`
}

type SubsnicResponseDto struct {
	Status  string        `json:"status"`
	Version string        `json:"version"`
	Error   SubsonicError `json:"error"`
}

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
	ArtistId int    `json:"artistId"`
	Artist   string `json:"artist"`
	AlbumId  string `json:"albumId"`
	Year     int    `json:"year"`
	Size     int    `json:"size"`
	Suffix   string `json:"suffix"`
	Duration int    `json:"duration"`
	Bitrate  int    `json:"bitRate"`
	Bitdepth int    `json:"bitDepth"`
}

type SubsonicTrack struct {
	Track
	Parent             string   `json:"parent"`
	IsDir              bool     `json:"isDir"`
	CovertArt          string   `json:"coverArt"`
	ContentType        string   `json:"contentType"`
	Suffix             string   `json:"suffix"`
	Path               string   `json:"path"`
	Created            string   `json:"created"`
	Type               string   `json:"type"`
	IsVideo            bool     `json:"isVideo"`
	Bpm                int      `json:"bpm"`
	Comment            string   `json:"comment"`
	SortName           string   `json:"sortName"`
	MediaType          string   `json:"mediaType"`
	MusicBrainzId      string   `json:"musicBrainzId"`
	Isrc               []any    `json:"isrc"`
	Genres             []string `json:"genres"`
	ReplayGain         any      `json:"replayGain"`
	ChannelCount       int      `json:"channelCount"`
	SamplingRate       int      `json:"samplingRate"`
	Moods              []any    `json:"moods"`
	Artists            []any    `json:"artists"`
	DisplayArtist      string   `json:"displayArtist"`
	AlbumArtists       []any    `json:"albumArtists"`
	DisplayAlbumArtist string   `json:"displayAlbumArtist"`
	Contributors       []any    `json:"contributors"`
	DisplayComposer    string   `json:"displayComposer"`
	ExplicitStatus     string   `json:"explicitStatus"`
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
