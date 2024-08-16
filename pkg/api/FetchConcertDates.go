package api

import (
	"encoding/json"
	"fmt"
	"groupie-Tracker/pkg/models"
	"net/http"
)

func FetchConcertDates() (models.ConcertDatesResponse, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error Failed to Fetch concert dates")
		return models.ConcertDatesResponse{}, fmt.Errorf("failed to fetch concert dates: %v", err)
	}
	defer resp.Body.Close()

	var concertDates models.ConcertDatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&concertDates); err != nil {
		fmt.Println("Failed to decode concert dates")
		return models.ConcertDatesResponse{}, fmt.Errorf("failed to decode concert dates: %v", err)
	}
	return concertDates, nil
}
