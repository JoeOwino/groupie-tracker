package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchAPI(api string) []byte {
	response, err := http.Get(api)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return body
}
