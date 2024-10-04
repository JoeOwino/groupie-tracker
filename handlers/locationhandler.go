package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Locationhandler(w http.ResponseWriter, r *http.Request) {
	locations := api.DecodeLocations()

	t, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, locations)
}