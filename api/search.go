package api

import (
	"strconv"
	"strings"
)

// SearchSuggestions returns a list of suggestions based on the search input.
func SearchResults(query string, artists []Artist, locations []Location) ([][]string, []Artist) {
	if len(query) < 2 {
		return nil, artists
	}

	suggestions := [][]string{}
	query = strings.ToLower(query)
	filteredArtists := []Artist{}

	for i, artist := range artists {
		isFound := false
		id := strconv.Itoa(artist.ArtistID)

		// Check for artist name
		if strings.Contains(strings.ToLower(artist.ArtistName), query) {
			suggestions = append(suggestions, []string{id, artist.ArtistName + " - artist/band"})
			isFound = true
		}

		// Check for band members
		for _, member := range artist.BandMembers {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, []string{id, member + " - member"})
				isFound = true
			}
		}
		// Check for first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			suggestions = append(suggestions, []string{id, artist.FirstAlbum + " - first album"})
			isFound = true
		}
		// Check for creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			suggestions = append(suggestions, []string{id, strconv.Itoa(artist.CreationDate) + " - creation date"})
			isFound = true
		}

		for _, locName := range locations[i].LocationName {
			if strings.Contains(strings.ToLower(locName), query) {
				suggestions = append(suggestions, []string{id, locName + " - location"})
				isFound = true
			}
		}

		if isFound {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return suggestions, filteredArtists
}
