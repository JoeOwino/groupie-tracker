package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Artisthandler(w http.ResponseWriter, r *http.Request) {
	artists := api.DecodeArtists()

	t, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, artists)
}

