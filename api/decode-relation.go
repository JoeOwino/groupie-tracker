package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type Relations struct {
	Index []Relation
}

type Relation struct {
	ArtistID  int      `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}

func DecodeRelations() Relations {
	apiBody := FetchAPI("https://groupietrackers.herokuapp.com/api/relation")
	Relations := Relations{}

	err := json.Unmarshal(apiBody, &Relations)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return Relations
}

func RelationMap(relations Relations) map[int]map[string][]string {
	rMap := make(map[int]map[string][]string)

	for _, index := range relations.Index {
		rMap[index.ArtistID] = index.Locations
	}

	return rMap
}