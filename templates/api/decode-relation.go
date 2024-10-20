package api

import (
	"encoding/json"
)

type Relations struct {
	Index []Relation
}

type Relation struct {
	ArtistID  int                 `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}

func DecodeRelations(RelationsAPI string) (Relations, error) {
	apiBody, err := FetchAPI(RelationsAPI)
	if err != nil {
		return Relations{}, err
	}
	relations := Relations{}

	err = json.Unmarshal(apiBody, &relations)
	if err != nil {
		return Relations{}, err
	}

	return relations, nil
}

func RelationMap(relations Relations) map[int]map[string][]string {
	rMap := make(map[int]map[string][]string)

	for _, index := range relations.Index {
		rMap[index.ArtistID] = index.Locations
	}

	return rMap
}
