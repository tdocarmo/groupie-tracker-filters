package loaders

import (
	"encoding/json"
	"io"
	"net/http"
)

// PageData is a struct representing data for rendering pages
type PageData struct {
	Data interface{} // Data to be displayed on the page
}

// GetDataFromURL retrieves data from an API URL, unmarshals it into the specified struct
func GetDataFromURL(url string, v interface{}) {
	response, err := http.Get(url)
	if err != nil {
		http.Error(nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		http.Error(nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
