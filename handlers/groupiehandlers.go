package handlers

import (
	"html/template"
	"net/http"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler(w, r, "templates/details.html")
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler(w, r, "templates/dates.html")
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler(w, r, "templates/locations.html")
}

func ConcertsHandler(w http.ResponseWriter, r *http.Request) {
	groupieHandler(w, r, "templates/concerts.html")
}

func groupieHandler(w http.ResponseWriter, r *http.Request, tpl string) {
	artist, err := Artist(r)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}

	if artist.ArtistName == "#Not Found" {
		errTxt := "Oops! The page you are looking for does not exist\n"
		ErrorHandler(w, r, http.StatusNotFound, errTxt, "404 Not Found", "Artists")
		return
	}

	t, err := template.ParseFiles(tpl)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error", "Error", "Reload")
		return
	}
	t.Execute(w, artist)
}
