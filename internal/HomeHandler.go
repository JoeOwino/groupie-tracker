package internal

import (
	"groupie-Tracker/pkg/api"
	"html/template"
	"net/http"
)

var tmplt, err = template.ParseGlob("./templates/*.html")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" && r.URL.Path != "/concert_dates" && r.URL.Path != "/locations" && r.URL.Path != "/relations" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with data
	err = tmplt.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
