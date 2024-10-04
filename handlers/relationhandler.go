package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

func Relationthandler(w http.ResponseWriter, r *http.Request) {
	relations := api.DecodeRelations()

	t, err := template.ParseFiles("templates/relations.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, relations)
}