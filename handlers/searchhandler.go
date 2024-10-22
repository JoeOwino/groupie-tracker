package handlers

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/api"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := api.DecodeArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	locations, err := api.DecodeLocations("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Ensure locations is a slice
	locationSlice := []api.Location{locations} // Wrap locations in a slice

	suggestions := api.SearchSuggestions(query, artists, locationSlice) // Use the slice

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
