package internal

import (
	"groupie-Tracker/pkg/api"
	"groupie-Tracker/pkg/models"
	"net/http"
	"strconv"
)

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	relations, err := api.FetchRelations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filteredRelations []models.Relation
	for _, relation := range relations.Index {
		if relation.Id == artistId {
			filteredRelations = append(filteredRelations, relation)
		}
	}
	// Execute the template with no data
	err = tmplt.ExecuteTemplate(w, "relations.html", filteredRelations)
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
