package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type Locations struct {
	Index []Location
}

type Location struct {
	ArtistID     int      `json:"id"`
	LocationName []string `json:"locations"`
}

func DecodeLocations() Locations {
	apiBody := FetchAPI("https://groupietrackers.herokuapp.com/api/locations")
	locations := Locations{}

	err := json.Unmarshal(apiBody, &locations)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return locations
}

func LocationMap(locations Locations) map[int][]string {
	lMap := make(map[int][]string)

	for _, Index := range locations.Index {
		lMap[Index.ArtistID] = Index.LocationName
	}

	return lMap
}
