package handlers

import (
	"net/http"
	"strconv"

	"groupie-tracker/api"
)

type ArtistDetails struct {
	ArtistID     int
	ArtistName   string
	ArtistImage  string
	BandMembers  []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Dates        []string
	Concerts     map[string][]string
}

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
func Artist(r *http.Request) (ArtistDetails, error) {
	id, err := GetID(r)
	if err != nil {
		return ArtistDetails{ArtistName: "#Not Found"}, nil
	}

	artistList, err := api.DecodeArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return ArtistDetails{}, err
	}

	ArtistsMap := api.ArtistsMap(artistList)
	artist, exist := ArtistsMap[id]
	if !exist {
		return ArtistDetails{ArtistName: "#Not Found"}, nil
	}

	strid := strconv.Itoa(id)

	location, err := api.DecodeLocation("https://groupietrackers.herokuapp.com/api/locations/" + strid)
	if err != nil {
		return ArtistDetails{}, err
	}

	date, err := api.DecodeDate("https://groupietrackers.herokuapp.com/api/dates/" + strid)
	if err != nil {
		return ArtistDetails{}, err
	}

	concert, err := api.DecodeRelations("https://groupietrackers.herokuapp.com/api/relation/" + strid)
	if err != nil {
		return ArtistDetails{}, err
	}

	artistDetails := ArtistDetails{
		ArtistID:     artist.ArtistID,
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Locations:    location.LocationName,
		Dates:        date.Date,
		Concerts:     concert.Locations,
	}

	return artistDetails, nil
}
