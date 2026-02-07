package utils

import (
	"fmt"
	"io"
	"net/http"
	dabtypes "nproxy/middlewares/dabmusic/types"
	libTypes "nproxy/server"
	"os"
	"strconv"

	"go.senan.xyz/taglib"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

type AlbumResult struct {
	libTypes.SubsonicResponseAlbumDto
	Song []libTypes.SubsonicTrack `json:"song"`
}
type AlbumResponse struct {
	libTypes.SubsonicResponse
	Album AlbumResult `json:"album"`
}
type DabToSubsonicAlbumResponse struct {
	SubsonicResponseDto AlbumResponse `json:"subsonic-response"`
}

// Necessary for `/rest/getAlbum`
func DabToSubsonicAlbum(
	album dabtypes.DabAlbum,
) DabToSubsonicAlbumResponse {
	subsonicResponse := libTypes.SubsonicResponse{
		Status:        "ok",
		Version:       "1.16.1",
		Type:          "navidrome",
		ServerVersion: "0.59.0 (cc3cca60)",
		OpenSubsonic:  true,
	}

	responseTracks := make([]libTypes.SubsonicTrack, 0)
	for _, track := range album.Tracks {
		responseTracks = append(responseTracks, libTypes.SubsonicTrack{
			Track: libTypes.Track{
				Id:       fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(int(track.Id))),
				Path:     "REMOTE",
				Title:    track.Title,
				Album:    track.Album,
				ArtistId: track.ArtistId,
				Artist:   track.Artist,
				AlbumId:  fmt.Sprintf("ext-dabmusic-%s", track.AlbumId),
				Year:     1978, // todo
				Size:     0,
				Suffix:   "TODO",
				Duration: track.Duration,
				Bitrate:  0,
				Bitdepth: 0,
			},
			Parent:             "TODO",
			IsDir:              false,
			CovertArt:          fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(int(track.Id))), // TODO
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

	response := DabToSubsonicAlbumResponse{
		SubsonicResponseDto: AlbumResponse{
			SubsonicResponse: subsonicResponse,
			Album: AlbumResult{
				SubsonicResponseAlbumDto: libTypes.SubsonicResponseAlbumDto{
					Id:            album.Id,
					Name:          album.Title,
					Artist:        album.Artist,
					ArtistId:      strconv.Itoa(responseTracks[0].ArtistId),
					CoverArt:      album.Cover,
					SongCount:     len(responseTracks),
					Duration:      0,
					PlayCount:     0,
					Created:       "",
					Year:          1978,
					Played:        "",
					UserRating:    1,
					Genres:        []string{},
					MusicBrainzId: "",
					IsCompilation: false,
					SortName:      album.Title,
					DiscTitles:    []string{},
					ReleaseDate:   "",
					ReleaseTypes:  []any{},
					RecordLabels:  []string{},
					Moods:         []string{},
					Artists:       []any{},
				},
				Song: responseTracks,
			},
		},
	}

	return response
}

type SearchResult struct {
	Song []libTypes.SubsonicTrack `json:"song"`
}
type Search3Response struct {
	libTypes.SubsonicResponse
	SearchResult3 SearchResult `json:"searchResult3"`
}
type Response struct {
	SubsonicResponseDto Search3Response `json:"subsonic-response"`
}

// Necessary for `/rest/search3` endpoint of subsonic
func DabToSubsonicTrack(
	tracks []dabtypes.DabTrack,
) Response {
	subsonicResponse := libTypes.SubsonicResponse{
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
				Id:       fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(int(track.Id))),
				Path:     "REMOTE",
				Title:    track.Title,
				Album:    track.Album,
				ArtistId: track.ArtistId,
				Artist:   track.Artist,
				AlbumId:  fmt.Sprintf("ext-dabmusic-%s", track.AlbumId),
				Year:     1978, // todo
				Size:     0,
				Suffix:   "TODO",
				Duration: track.Duration,
				Bitrate:  0,
				Bitdepth: 0,
			},
			Parent:             "TODO",
			IsDir:              false,
			CovertArt:          fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(int(track.Id))), // TODO
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

	response := Response{
		SubsonicResponseDto: Search3Response{
			SubsonicResponse: subsonicResponse,
			SearchResult3: SearchResult{
				Song: responseTracks,
			},
		},
	}

	return response
}

// Necessary for `/api/song` endpoint of navidrome
func DabToNavidromeTrack(
	track dabtypes.DabTrack,
) libTypes.NavidromeTrack {
	return libTypes.NavidromeTrack{
		Track: libTypes.Track{
			Id:       fmt.Sprintf("ext-%s-%s", "dabmusic", strconv.Itoa(int(track.Id))),
			Path:     "REMOTE",
			Title:    track.Title,
			Album:    track.Album,
			ArtistId: track.ArtistId,
			Artist:   track.Artist,
			AlbumId:  fmt.Sprintf("ext-dabmusic-%s", track.AlbumId),
			Year:     1978,
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

func AddMetadata(targetFile string, track *dabtypes.DabTrack) error {
	// res, err := _request(metadatas.Cover, false, []QueryParams{})
	req, err := http.NewRequest("GET", track.Cover, nil)
	if err != nil {
		return fmt.Errorf("can't download cover")
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("can't download cover")
	}
	defer res.Body.Close()

	coverBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = taglib.WriteTags(targetFile, map[string][]string{
		taglib.Title:  {track.Title},
		taglib.Artist: {track.Artist},
		taglib.Album:  {track.Album},
		taglib.Date:   {track.ReleaseDate},
	}, 0)

	if err != nil {
		return fmt.Errorf("unable to write metadata to track")
	}

	if len(coverBytes) > 0 {
		err = taglib.WriteImage(targetFile, coverBytes)
	}

	if err != nil {
		return fmt.Errorf("unable to picture meadata to track")
	}

	return nil
}
