package api

import (
	"encoding/json"
	"fmt"
	"groupie-Tracker/pkg/models"
	"net/http"
)

func FetchLocations() ([]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	var wrapper models.LocationWrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode locations: %v", err)
	}

	return wrapper.Index, nil
}
