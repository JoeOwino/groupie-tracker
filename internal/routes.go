package internal

import (
	"net/http"
)

func Routes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/relations", RelationHandler)
	http.HandleFunc("/locations", LocationHandler)
	http.HandleFunc("/concert_dates/", ConcertDatesHandler)
	http.HandleFunc("/artistProfile", ArtistProfileHandler)
	http.HandleFunc("/search", SearchHandler)


}
