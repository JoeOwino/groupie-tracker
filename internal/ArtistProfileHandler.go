package internal

import (
	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
	"net/http"
	"strconv"
)

func ArtistProfileHandler(w http.ResponseWriter, r *http.Request) {

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

	artists, err := api.FetchArtists()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var selectedArtist models.Artist

	for _, artist := range artists {
		if artist.Id == artistId {
			selectedArtist = artist
			break
		}
	}

	if selectedArtist.Id == 0 {
		ErrorPage(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Fetch location data
	locations, err := api.FetchLocations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch concert dates
	concertDates, err := api.FetchConcertDates()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch relations
	relations, err := api.FetchRelations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter data by artist ID
	var artistLocations []models.Location
	for _, location := range locations {
		if location.Id == artistId {
			artistLocations = append(artistLocations, location)
		}
	}

	var artistConcertDates []models.ConcertDate
	for _, concertDate := range concertDates.Index {
		if concertDate.Id == artistId {
			artistConcertDates = append(artistConcertDates, concertDate)
		}
	}

	var artistRelations []models.Relation
	for _, relation := range relations.Index {
		if relation.Id == artistId {
			artistRelations = append(artistRelations, relation)
		}
	}

	// Combine artist data with related data
	data := struct {
		Artist       models.Artist
		Locations    []models.Location
		ConcertDates []models.ConcertDate
		Relations    []models.Relation
	}{
		Artist:       selectedArtist,
		Locations:    artistLocations,
		ConcertDates: artistConcertDates,
		Relations:    artistRelations,
	}

	// Execute the template with data
	err = tmplt.ExecuteTemplate(w, "artistProfile.html", data)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
