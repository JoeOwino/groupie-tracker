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

func LocationMap(locations Locations) map[int][]string {
	lMap := make(map[int][]string)

	for _, Index := range locations.Index {
		lMap[Index.ArtistID] = Index.LocationName
	}

	return lMap
}
