package internal

import (
	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
	"net/http"
	"strconv"
)

func ConcertDatesHandler(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}
	concertDates, err := api.FetchConcertDates()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredDates := []models.ConcertDate{}
	for _, date := range concertDates.Index {
		if date.Id == artistId {
			filteredDates = append(filteredDates, date)
		}
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "concert_dates.html", filteredDates)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
