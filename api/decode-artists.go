package api

import (
	"encoding/json"
)

// Artist struct represents the structure for storing artist-related data.
// Fields are mapped to the corresponding JSON keys using struct tags.
type Artist struct {
	ArtistID     int      `json:"id"`
	ArtistName   string   `json:"name"`
	ArtistImage  string   `json:"image"`
	BandMembers  []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// DecodeArtists fetches and decodes the artist data from the given API.
func DecodeArtists(artistAPI string) ([]Artist, error) {
	apiBody, err := FetchAPI(artistAPI)
	if err != nil {
		return nil, err
	}
	// Initialize a slice of Artist to hold the decoded data
	artists := []Artist{}

	err = json.Unmarshal(apiBody, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
