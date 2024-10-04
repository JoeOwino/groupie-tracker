package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type Artist struct {
	ArtistID		int			`json:"id"`
	ArtistName		string		`json:"name"`
	ArtistImage		string		`json:"image"`
	BandMembers		[]string	`json:"members"`
	CreationDate	int			`json:"creationDate"`
	FirstAlbum		string		`json:"firstAlbum"`
	LocationAPI		string		`json:"locations"`
	ConcertDatesAPI	string		`json:"concertDates"`
	RelationsAPI	string		`json:"relation"`
} 

func DecodeArtists() []Artist {
	apiBody := FetchAPI("https://groupietrackers.herokuapp.com/api/artists")
	artists := []Artist{}

	err := json.Unmarshal(apiBody, &artists)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return artists
}
