package handlers

import (
	"net/http"
	"strconv"

	"groupie-tracker/api"
)

// GetID retrieves the artist ID from the HTTP request URL query parameters.
// It returns the artist ID as an integer and an error if any occurs during conversion.
func GetID(r *http.Request) (int, error) {
	id := r.URL.Query().Get("id")

	artistID, err := strconv.Atoi(id)
	if err != nil || artistID < 1 {
		return artistID, err
	}

	return artistID, nil
}

// Artist retrieves detailed information about an artist based on the provided request.
func Artist(r *http.Request) (api.ArtistsList, error) {
	id, err := GetID(r)
	if err != nil {
		return api.ArtistsList{ArtistName: "#Not Found"}, nil
	}

	artistList, err := api.ArtistMap()
	if err != nil {
		return api.ArtistsList{}, err
	}

	artist, exist := artistList[id]
	if !exist {
		return api.ArtistsList{ArtistName: "#Not Found"}, nil
	}

	return artist, nil
}
