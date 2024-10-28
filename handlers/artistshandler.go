package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

type ArtistData struct {
	Data  []api.Artist
	Query string
}

// Artistshandler handles the HTTP requests for the artists' page.
// It responds to GET requests by fetching artist data and rendering the artists.html template.
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.DecodeArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	locations, err := api.DecodeLocations("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	query := r.FormValue("search_query")
	_, filteredArtists := api.SearchResults(query, artists, locations.Index)

	data := ArtistData{
		Data:  filteredArtists,
		Query: query,
	}

	t, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}
	t.Execute(w, data)
}
