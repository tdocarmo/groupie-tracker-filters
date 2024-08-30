package loaders

import (
	"groupie-tracker-filters/constants"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var artists []constants.Artist
var locations []constants.Locations

// Index is a handler function to load data for the index page
func Index(w http.ResponseWriter, r *http.Request, t *template.Template, members []string) {
	// Fetch data from API for all artists
	GetDataFromURL(constants.API_ARTISTS_URL, &artists)
	done := make(chan bool, 1)
	go RécupérationLieux(done)
	// Supprimer les duplicas.

	// Parse filter criteria from query parameters
	criteria := parseFilterCriteria(r, members)

	// Filter the data based on the criteria
	filteredArtists := filterData(artists, criteria)

	pageData := PageData{Data: filteredArtists}
	t.Execute(w, pageData)
}

// parseFilterCriteria parses the filter criteria from query parameters
func parseFilterCriteria(r *http.Request, members []string) constants.FilterCriteria {
	minCreationDate := r.URL.Query().Get("minCreationDate")
	maxCreationDate := r.URL.Query().Get("maxCreationDate")
	minFirstAlbumDate := r.URL.Query().Get("minFirstAlbumDate")
	maxFirstAlbumDate := r.URL.Query().Get("maxFirstAlbumDate")
	members = r.URL.Query()["members"]
	concertLocations := r.URL.Query().Get("concertLocations")

	// Decode the concertLocations parameter
	//decodedLocations, _ := url.QueryUnescape(concertLocations)

	// Split the locations by comma
	//locationList := strings.Split(decodedLocations, ",")

	// Conversion des valeurs des cases cochées en entiers
	var membersInt []int
	for _, m := range members {
		memberInt, _ := strconv.Atoi(m)
		membersInt = append(membersInt, memberInt)
	}

	return constants.FilterCriteria{
		MinCreationDate:   minCreationDate,
		MaxCreationDate:   maxCreationDate,
		MinFirstAlbumDate: minFirstAlbumDate,
		MaxFirstAlbumDate: maxFirstAlbumDate,
		Members:           membersInt,
		ConcertLocations:  concertLocations,
	}
}

func RécupérationLieux(done chan bool) {
	// Associer les localisations à chaque artiste
	for i := range artists {
		var location constants.Locations
		id := "/" + strconv.Itoa(i+1)
		GetDataFromURL(constants.API_LOCATIONS_URL+id, &location)
		locations = append(locations, location)
		// Concaténer les localisations en une chaîne pour faciliter le filtrage
		artists[i].Locations = strings.Join(location.Locations, ",")
	}
	done <- true
}
