package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

type ArtistDetails struct {
	ArtistName   string
	ArtistImage  string
	BandMembers  []string
	CreationDate int
	FirstAlbum   string
	Concerts     map[string][]string
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	artist := Artist(2)

	t, err := template.ParseFiles("templates/details.html")
	if err != nil {
		fmt.Println("err")
	}
	t.Execute(w, artist)
}

func Artist(id int) ArtistDetails {
	artistList := api.DecodeArtists()
	artist := artistList[id-1]

	concerts := api.RelationMap(api.DecodeRelations())

	artistD := ArtistDetails{
		ArtistName:   artist.ArtistName,
		ArtistImage:  artist.ArtistImage,
		BandMembers:  artist.BandMembers,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Concerts:     concerts[id],
	}

	return artistD
}
