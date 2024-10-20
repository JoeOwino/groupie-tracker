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

	if len(artistList) < id {
		return ArtistDetails{ArtistName: "#Not Found"}, nil
	}

	artist := artistList[id-1]

	concertAPI, errConst := api.DecodeRelations("https://groupietrackers.herokuapp.com/api/relation")
	dateAPI, errDt := api.DecodeDates("https://groupietrackers.herokuapp.com/api/dates")
	locationAPI, errLoc := api.DecodeLocations("https://groupietrackers.herokuapp.com/api/locations")
	if errConst != nil || errDt != nil || errLoc != nil {
		return ArtistDetails{}, err
	}

	concerts := api.RelationMap(concertAPI)
	dates := api.DateMap(dateAPI)
	locations := api.LocationMap(locationAPI)

	artistD := ArtistDetails{
		ArtistID:     id,
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Dates:        dates[id],
		Locations:    locations[id],
		Concerts:     concerts[id],
	}

	return artistD, nil
}
