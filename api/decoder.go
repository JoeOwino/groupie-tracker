package api

import "strconv"

type ArtistsList struct {
	ArtistName   string
	ArtistImage  string
	BandMembers  []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Dates        []string
	Concerts     map[string][]string
}

func ArtistMap() (map[int]ArtistsList, error) {
	artistMap := map[int]ArtistsList{}

	artists, err := DecodeArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return artistMap, err
	}

	for _, artist := range artists {
		id := artist.ArtistID
		locationAPI := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id)
		location, err := DecodeLocations(locationAPI)
		if err != nil {
			return artistMap, err
		}

		dateAPI := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id)
		date, err := DecodeDates(dateAPI)
		if err != nil {
			return artistMap, err
		}

		concertAPI := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)
		concert, err := DecodeRelations(concertAPI)
		if err != nil {
			return artistMap, err
		}

		band := ArtistsList{
			ArtistName:   artist.ArtistName,
			ArtistImage:  artist.ArtistImage,
			BandMembers:  artist.BandMembers,
			CreationDate: artist.CreationDate,
			FirstAlbum:   artist.ArtistName,
			Locations:    location.LocationName,
			Dates:        date.Date,
			Concerts:     concert.Locations,
		}

		artistMap[id] = band
	}

	return artistMap, nil
}
