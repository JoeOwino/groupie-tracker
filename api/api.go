package api

import (
	"io"
	"net/http"
)

// Sends a GET request to the provided API endpoint and returns response as a byte of slice
func FetchAPI(api string) ([]byte, error) {
	response, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
