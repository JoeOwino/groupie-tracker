package handlers

import (
	"html/template"
	"net/http"
)

// Holds data for error handling in the application.
type Error struct {
	ErrNo     int
	ErrText   string
	ErrTitle  string
	ErrButton string
}

// Handles rendering of error page
func ErrorHandler(w http.ResponseWriter, r *http.Request, errNo int, errTxt, errTitle, errBtn string) {
	errData := Error{ErrNo: errNo, ErrText: errTxt, ErrTitle: errTitle, ErrButton: errBtn}

	tmp, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.WriteHeader(errNo)
	tmp.Execute(w, errData)
}
