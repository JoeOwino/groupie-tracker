package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"groupie-tracker/api"
)

func TestArtist(t *testing.T) {
	// Mock data for artists, locations, and relations
	artistsMock := []api.Artist{
		{
			ArtistID:     1,
			ArtistImage:  "image_url",
			ArtistName:   "Band 1",
			BandMembers:  []string{"Member 1", "Member 2"},
			CreationDate: 2000,
			FirstAlbum:   "Album 1",
		},
	}

	artistsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(artistsMock)
		_, err := Artist(r)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	}))
	defer artistsServer.Close()
}

func TestPathHandler(t *testing.T) {
	defer changeToParentDir(t)()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Expected no error found %v", err)
	}

	rr := httptest.NewRecorder()
	artistsHandler := http.HandlerFunc(ArtistsHandler)
	detailsHandler := http.HandlerFunc(DetailsHandler)
	concertsHandler := http.HandlerFunc(ConcertsHandler)

	artistsHandler.ServeHTTP(rr, req)
	detailsHandler.ServeHTTP(rr, req)
	concertsHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Helper function to temporarily change to parent directory
func changeToParentDir(t *testing.T) func() {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir("../")
	if err != nil {
		t.Fatal(err)
	}
	return func() {
		err := os.Chdir(originalDir)
		if err != nil {
			t.Fatal(err)
		}
	}
}
