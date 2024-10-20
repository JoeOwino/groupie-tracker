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

	locationsMock := Location{
		ArtistID: 1, LocationName: []string{"City 1", "City 2"},
	}

	dateMock := Date{
		ArtistID: 1, Date: []string{"City 1", "City 2"},
	}

	relationsMock := Relation{
			ArtistID: 1, Locations: map[string][]string{"2024": {"City 1", "City 2"}},
	}

	// Create mock servers
	artistsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(artistsMock)
	}))
	defer artistsServer.Close()

	locationsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(locationsMock)
	}))
	defer locationsServer.Close()

	datesServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(dateMock)
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


	if artistsMock[0].ArtistName != artists[0].ArtistName {
		t.Errorf("Expected %v got %v", artistsMock[0].ArtistName, artists[0].ArtistName)
	}

	// if lMap == nil {
	// 	t.Errorf("Expected Location map got nil")
	// }

	// if dMap == nil {
	// 	t.Errorf("Expected Date map got nil")
	// }

	// if rMap == nil {
	// 	t.Errorf("Expected Date map got nil")
	// }
}
