package routes

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"groupie-tracker-filters/loaders"
	"strconv"
	"strings"
)

// artist registers routes related to artists
func artist() {
	http.HandleFunc("/artist/", ArtistHandler)
}

// ArtistHandler is the handler for the /artist/ route
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect if the path is "/artist/"
	if r.URL.Path == "/artist/" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	// Extract artist ID from the URL path
	id, err := GetArtistsFromURL(r.URL.Path)
	if err != nil {
		http.RedirectHandler("/error500", http.StatusTemporaryRedirect)
		return
	}

	// Convert the ID to an integer
	str, err1 := strconv.Atoi(id)
	if err1 != nil {
		fmt.Println("Error 2")
		http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
		return
	}

	// Validate the artist ID range
	if str > 52 || str < 1 {
		fmt.Println("Error: Invalid artist ID")
		http.Redirect(w, r, "/error400", http.StatusTemporaryRedirect)
		return
	} else {
		// Load the artist page template and data
		pageTemplate, err := template.ParseFiles("client/artist.html")
		if err != nil {
			fmt.Println("Error 2")
			http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
			return
		}
		loaders.Artist(w, r, id, pageTemplate)
	}
}

// GetArtistsFromURL extracts the artist ID from the URL
func GetArtistsFromURL(url string) (string, error) {
	// Check if the URL contains "/artist/"
	if !(strings.Contains(url, "/artist/")) {
		return "", errors.New("URL is not an artist URL")
	}

	// Extract the artist ID from the URL
	id, found := strings.CutPrefix(url, "/artist/")
	if !found {
		fmt.Println("Error 3")
		return "", errors.New("URL is not an artist URL")
	}
	return id, nil
}
