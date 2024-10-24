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

	suggestions, _ := api.SearchResults(query, artists, locations.Index)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
