package utils

import (
	"fmt"
	libTypes "nproxy/lib"
	dabtypes "nproxy/middlewares/dabmusic/types"
	"strconv"
)

type SearchResult struct {
	Song []libTypes.SubsonicTrack `json:"song"`
}

type SubsonicResponse struct {
	Status        string       `json:"status"`
	Version       string       `json:"version"`
	Type          string       `json:"type"`
	ServerVersion string       `json:"serverVersion"`
	OpenSubsonic  bool         `json:"openSubsonic"`
	SearchResult3 SearchResult `json:"searchResult3"`
}

type Response struct {
	SubsonicResponseDto SubsonicResponse `json:"subsonic-response"`
}

// Necessary for `/rest/search3` endpoint of subsonic
func DabToSubsonicTrack(
	// {
	//     "subsonic-response": {
	//         "status": "ok",
	//         "version": "1.16.1",
	//         "type": "navidrome",
	//         "serverVersion": "0.59.0 (cc3cca60)",
	//         "openSubsonic": true,
	//         "searchResult3": {
	//             "song": [
	tracks []dabtypes.DabTrack,
) Response {

	subResponse := Response{}

	response := SubsonicResponse{
		Status:        "ok",
		Version:       "1.16.1",
		Type:          "navidrome",
		ServerVersion: "0.59.0 (cc3cca60)",
		OpenSubsonic:  true,
	}

	responseTracks := make([]libTypes.SubsonicTrack, 0)
	for _, track := range tracks {
		responseTracks = append(responseTracks, libTypes.SubsonicTrack{
			Track: libTypes.Track{
				Id:       fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(track.Id)),
				Path:     "REMOTE",
				Title:    track.Title,
				Album:    track.Album,
				ArtistId: "TODO",
				Artist:   track.Artist,
				AlbumId:  "TODO",
				Year:     track.ReleaseDate,
				Size:     0,
				Suffix:   "TODO",
				Duration: track.Duration,
				Bitrate:  0,
				Bitdepth: 0,
			},
			Parent:             "TODO",
			IsDir:              false,
			CovertArt:          fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(track.Id)), // TODO
			ContentType:        "audio/mpeg",
			Suffix:             "mp3",
			Path:               "TODO",
			Created:            track.ReleaseDate,
			Type:               "music",
			IsVideo:            false,
			Bpm:                0,
			Comment:            "",
			SortName:           track.Title,
			MediaType:          "song",
			MusicBrainzId:      "",
			Isrc:               []any{},
			Genres:             []string{},
			ReplayGain:         "",
			ChannelCount:       2,
			SamplingRate:       44100,
			Moods:              []any{},
			Artists:            []any{},
			DisplayArtist:      track.Artist,
			AlbumArtists:       []any{},
			DisplayAlbumArtist: track.Artist,
			Contributors:       []any{},
			DisplayComposer:    "",
			ExplicitStatus:     "",
		})
	}

	response.SearchResult3.Song = responseTracks
	subResponse.SubsonicResponseDto = response

	return subResponse
}

// Necessary for `/api/song` endpoint of navidrome
func DabToNavidromeTrack(
	track dabtypes.DabTrack,
) libTypes.NavidromeTrack {
	return libTypes.NavidromeTrack{
		Track: libTypes.Track{
			Id:       fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(track.Id)),
			Path:     "REMOTE",
			Title:    track.Title,
			Album:    track.Album,
			ArtistId: "TODO",
			Artist:   track.Artist,
			AlbumId:  "TODO",
			Year:     track.ReleaseDate,
			Size:     0,
			Suffix:   "TODO",
			Duration: track.Duration,
			Bitrate:  0,
			Bitdepth: 0,
		},
		Bookmarkposition:     0,
		Libraryid:            1,
		LibraryPath:          "/music",
		Libraryname:          "Music Library",
		Folderid:             "REMOTE_TODO?",
		Albumartist:          track.Artist,
		HasCoverArt:          true,
		Tracknumber:          track.TrackNumber,
		Discnumber:           1,
		Date:                 0,
		Samplerate:           44100,
		Channels:             2,
		Genre:                "",
		Ordertitle:           track.Title,
		Orderalbumname:       track.Album,
		Orderartistname:      track.Artist,
		Orderalbumartistname: track.Artist,
		Compilation:          false,
		Lyrics:               "[]",
		Explicitstatus:       "",
		Rgalbumgain:          nil,
		Rgalbumpeak:          nil,
		Rgtrackgain:          nil,
		Rgtrackpeak:          nil,
		Participants: libTypes.TrackParticipants{
			AlbumArtist: []libTypes.TrackAlbumArtist{
				libTypes.TrackAlbumArtist{
					Id:      "",
					Name:    track.Artist,
					Missing: false,
				},
			},
			Artist: []libTypes.TrackAlbumArtist{
				libTypes.TrackAlbumArtist{
					Id:      "",
					Name:    track.Artist,
					Missing: false,
				},
			},
		},
		Missing:   false,
		Birthtime: track.ReleaseDate,
		Createdat: track.ReleaseDate,
		Updatedat: track.ReleaseDate,
	}
}
