package models

type ConcertDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type ConcertDatesResponse struct {
	Index []ConcertDate `json:"index"`
}
