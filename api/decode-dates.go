package api

import (
	"encoding/json"
)

// Dates struct contains a slice of Date struct, which holds concert dates for multiple artists.
type Dates struct {
	Index []Date
}

// Date struct holds the ArtistID and a slice of strings representing concert dates.
type Date struct {
	ArtistID int      `json:"id"`
	Date     []string `json:"dates"`
}

// DecodeDates fetches and decodes the dates data from the API.
func DecodeDate(dateAPI string) (Date, error) {
	apiBody, err := FetchAPI(dateAPI)
	if err != nil {
		return Date{}, err
	}

	dates := Date{}
	err = json.Unmarshal(apiBody, &dates)
	if err != nil {
		return Date{}, err
	}

	return dates, nil
}

func DecodeDates(datesAPI string) (Dates, error) {
	apiBody, err := FetchAPI(datesAPI)
	if err != nil {
		return Dates{}, err
	}

	dates := Dates{}
	err = json.Unmarshal(apiBody, &dates)
	if err != nil {
		return Dates{}, err
	}

	return dates, nil
}
