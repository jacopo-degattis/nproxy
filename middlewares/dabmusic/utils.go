package dabmusic

import (
	"fmt"
	libTypes "nproxy/lib"
	"strconv"
)

func DabToNavidromeTrack(
	track DabTrack,
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
