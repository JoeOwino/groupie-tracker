package api

import (
	"strings"
	"strconv"
)

type SearchedArtists struct {
	ArtistID     int
	ArtistName   []string
	BandMembers  []string
	CreationDate []int
	FirstAlbum   []string
	Locations    []string
	Dates        []string
}

// SearchSuggestions returns a list of suggestions based on the search input.
func SearchSuggestions(query string, artists []Artist, locations []Location) []string {
	var suggestions []string
	query = strings.ToLower(query)

	// Create a map for faster lookups
	// artistMap := make(map[string]Artist)
	// for _, artist := range artists {
	// 	artistMap[artist.ArtistName] = artist
	// }

	for _, artist := range artists {
		// Check for artist name
		if strings.Contains(strings.ToLower(artist.ArtistName), query) {
			suggestions = append(suggestions, artist.ArtistName+" - artist/band")
		}
		// Check for band members
		for _, member := range artist.BandMembers {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, member+" - member")
			}
		}
		// Check for first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			suggestions = append(suggestions, artist.FirstAlbum+" - first album")
		}
		// Check for creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			suggestions = append(suggestions, strconv.Itoa(artist.CreationDate)+" - creation date")
		}
	}

	// Check for locations
	for _, location := range locations {
		for _, locName := range location.LocationName {
			if strings.Contains(strings.ToLower(locName), query) {
				suggestions = append(suggestions, locName+" - location")
			}
		}
	}

	return suggestions
}
