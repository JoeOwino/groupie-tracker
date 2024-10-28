package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockArtistAPI = `{"ArtistID": 1,
 					"ArtistName": "Test Artist",
					"ArtistImage": "Image url",
					"BandMembers": ["Member1", "Member2", "Member3"],
					"CreationDate": "1/1/2024",
					"FirstAlbum": "1/2/2024"}`

// Helper function to create a mock server for dates API
func MockerServer(mockResponse string) *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, mockResponse)
	}))
	return mockServer
}

func TestFetchAPI(t *testing.T) {
	// mockResponse := `{"message": "hello world"}`
	server := MockerServer(mockArtistAPI)

	defer server.Close()

	result, err := FetchAPI(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Verify the result
	expected := mockArtistAPI
	if string(result) != expected {
		t.Errorf("expected %v, got %v", expected, string(result))
	}
}

func TestDecoders(t *testing.T) {
	// Mock data for artists, locations, and relations
	artistsMock := []Artist{
		{
			ArtistID:     1,
			ArtistImage:  "image_url",
			ArtistName:   "Band 1",
			BandMembers:  []string{"Member 1", "Member 2"},
			CreationDate: 2000,
			FirstAlbum:   "Album 1",
		},
	}

	locationMock := Location{
		ArtistID: 1, LocationName: []string{"City 1", "City 2"},
	}

	locationsMock := Locations{
		Index: []Location{locationMock},
	}

	dateMock := Date{
		ArtistID: 1, Date: []string{"City 1", "City 2"},
	}

	datesMock := Dates{
		Index: []Date{dateMock},
	}

	relationsMock := Relation{
		ArtistID: 1, Locations: map[string][]string{"2024": {"City 1", "City 2"}},
	}

	// Create mock servers
	artistsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(artistsMock)
	}))
	defer artistsServer.Close()

	locationServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(locationMock)
	}))
	defer locationServer.Close()

	locationsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(locationsMock)
	}))
	defer locationsServer.Close()

	dateServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(dateMock)
	}))
	defer dateServer.Close()

	datesServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(datesMock)
	}))
	defer datesServer.Close()

	relationsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(relationsMock)
	}))
	defer relationsServer.Close()

	// Call the function to test
	artists, err := DecodeArtists(artistsServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	location, err := DecodeLocation(locationServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	date, err := DecodeDate(dateServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	relation, err := DecodeRelations(dateServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	locations, err := DecodeLocations(locationsServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	dates, err := DecodeDates(datesServer.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	artistMap := ArtistsMap(artists)

	if artistsMock[0].ArtistName != artists[0].ArtistName {
		t.Errorf("Expected %v got %v", artistsMock[0].ArtistName, artists[0].ArtistName)
	}

	if locationMock.LocationName[0] != location.LocationName[0] {
		t.Errorf("Expected %v got %v", locationMock.LocationName[0], location.LocationName[0])
	}

	if dateMock.Date[0] != date.Date[0] {
		t.Errorf("Expected %v got %v", dateMock.Date[0], date.Date[0])
	}

	if relationsMock.ArtistID != relation.ArtistID {
		t.Errorf("Expected %v got %v", relationsMock.ArtistID, relation.ArtistID)
	}

	if locations.Index == nil {
		t.Errorf("Expected Locations map got nil")
	}

	if dates.Index == nil {
		t.Errorf("Expected Dates map got nil")
	}

	if artistMap == nil {
		t.Errorf("Expected Artits map got nil")
	}

	suggestions, filteredArtist := SearchResults("Band", artists, locations.Index)

	if suggestions == nil {
		t.Errorf("Expected Suggestions map got nil")
	}

	if filteredArtist == nil {
		t.Errorf("Expected Artists map got nil")
	}
}
