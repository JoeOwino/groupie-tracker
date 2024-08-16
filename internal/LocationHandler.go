package internal

import (
	"fmt"
	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
	"net/http"
	"strconv"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("location error")
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	locations, err := api.FetchLocations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredLocations []models.Location
	for _, location := range locations {
		if location.Id == artistId {
			filteredLocations = append(filteredLocations, location)
		}
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "locations.html", filteredLocations)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
