package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"nproxy/lib"

	dabtypes "nproxy/middlewares/dabmusic/types"
)

// client to interface with dabmsusic api

type DabClient struct {
	BaseUrl string
}

// TODO: improve this function to work as a general fetch
// now is only for search with `q` query parameter!
func (d *DabClient) Search(query string, queryType string) ([]dabtypes.DabTrack, error) {
	type SearchResponse struct {
		Tracks []dabtypes.DabTrack `json:"tracks"`
	}

	encodedQuery := url.QueryEscape(query)

	fullPath := fmt.Sprintf(
		"%s/api/search?q=%s",
		d.BaseUrl,
		encodedQuery,
	)

	res, err := lib.Fetch(
		fullPath,
		"GET",
		nil,
		nil,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch dabmusic endpoint %s", "/api/search")
	}

	var response SearchResponse
	err = json.NewDecoder(bytes.NewReader(res)).Decode(&response)

	if err != nil {
		return nil, fmt.Errorf("unable to decode dabmusic response for endpoint %s", "/api/search")
	}

	return response.Tracks, nil
}

func (d *DabClient) GetTrackMetadata(trackId string) (*dabtypes.DabTrack, error) {
	tracks, err := d.Search(trackId, "track")

	if err != nil {
		return nil, fmt.Errorf("error while fetching /api/search: %s", err.Error())
	}

	// TODO: do i want to return an error if no track metadata is found for given id?
	if len(tracks) == 0 {
		return nil, fmt.Errorf("no tracks found for id %s", trackId)
	}

	return &tracks[0], nil
}

func (d *DabClient) GetTrackStreamUrl(trackId string) (*string, error) {
	type StreamResponse struct {
		Url string `json:"url"`
	}

	fullPath := fmt.Sprintf(
		"%s/api/stream?trackId=%s&quality=5",
		d.BaseUrl,
		trackId,
	)

	res, err := lib.Fetch(
		fullPath,
		"GET",
		nil,
		nil,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch dabmusic endpoint %s", "/api/stream")
	}

	var response StreamResponse
	err = json.NewDecoder(bytes.NewReader(res)).Decode(&response)

	if err != nil {
		return nil, fmt.Errorf("unable to decode dabmusic response for endpoint %s", "/api/stream")
	}

	return &response.Url, nil
}
