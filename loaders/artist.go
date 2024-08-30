package loaders

import (
	"fmt"
	"html/template"
	"net/http"
	"groupie-tracker-filters/constants"
)

// artistData is a struct representing combined data for an artist
type artistData struct {
	Artist    constants.Artist    // Information about the artist
	Relations constants.Relations // Information about the artist's relations
	Dates     constants.Dates     // Information about the artist's dates
	Locations constants.Locations // Information about the artist's locations
}

// Artist is a handler function to load data for a specific artist
func Artist(w http.ResponseWriter, r *http.Request, id string, t *template.Template) {
	var fullData artistData

	// Fetch data from API for the specified artist ID
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_ARTISTS_URL, id), &fullData.Artist)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_RELATION_URL, id), &fullData.Relations)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_DATES_URL, id), &fullData.Dates)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_LOCATIONS_URL, id), &fullData.Locations)

	pageData := PageData{Data: fullData}
	t.Execute(w, pageData)
}
