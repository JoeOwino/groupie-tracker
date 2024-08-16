package internal

import (
	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("query")
	if query == "" {
		ErrorPage(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []models.Artist
	// Example: Filter from a predefined list
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}

	if err = tmplt.ExecuteTemplate(w, "search_result.html", results); err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
