package api

import (
	"encoding/json"
)

type Relation struct {
	ArtistID  int                 `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}

func DecodeRelations(RelationAPI string) (Relation, error) {
	apiBody, err := FetchAPI(RelationAPI)
	if err != nil {
		return Relation{}, err
	}
	relations := Relation{}

	err = json.Unmarshal(apiBody, &relations)
	if err != nil {
		return Relation{}, err
	}

	return relations, nil
}
