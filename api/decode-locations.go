package api

import (
	"encoding/json"
)

type Locations struct {
	Index []Location
}

type Location struct {
	ArtistID     int      `json:"id"`
	LocationName []string `json:"locations"`
}

func DecodeLocation(locationAPI string) (Location, error) {
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

func DecodeLocations(locationAPI string) (Locations, error) {
	apiBody, err := FetchAPI(locationAPI)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}

	err = json.Unmarshal(apiBody, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
