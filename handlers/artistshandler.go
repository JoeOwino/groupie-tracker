package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

// Artistshandler handles the HTTP requests for the artists' page.
// It responds to GET requests by fetching artist data and rendering the artists.html template.
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.DecodeArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	t, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}
	t.Execute(w, artists)
}
