package api

import (
	"encoding/json"
)

type Location struct {
	ArtistID     int      `json:"id"`
	LocationName []string `json:"locations"`
}

func DecodeLocations(locationAPI string) (Location, error) {
	apiBody, err := FetchAPI(locationAPI)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}

	err = json.Unmarshal(apiBody, &locations)
	if err != nil {
		return Location{}, err
	}

	return locations, nil
}
