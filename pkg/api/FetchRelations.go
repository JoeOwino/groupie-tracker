package api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"groupie-Tracker/pkg/models"
)

func FetchRelations() (models.RelationsResponse, error) {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil {
        return models.RelationsResponse{}, fmt.Errorf("failed to fetch relations: %v", err)
    }
    defer resp.Body.Close()

    var relations models.RelationsResponse
    if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
        return models.RelationsResponse{}, fmt.Errorf("failed to decode relations: %v", err)
    }
    return relations, nil
}
