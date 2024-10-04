package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Datehandler(w http.ResponseWriter, r *http.Request) {
	dates := api.DecodeDates()

	t, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, dates)
}