package constants

// Artist represents the structure of an artist in the Groupie Trackers API
type Artist struct {
	ID           uint     // Artist ID
	Image        string   // URL for the artist's image
	Name         string   // Artist name
	Members      []string // List of members in the band
	CreationDate int      // Year of the band's creation
	FirstAlbum   string   // Title of the first album
	Locations    string   // URL for artist's locations
	ConcertDates string   // URL for artist's concert dates
	Relations    string   // URL for artist's relations
}

// Locations represents the structure of locations in the Groupie Trackers API
type Locations struct {
	ID        uint     `json:"id"`        // Locations ID
	Locations []string `json:"locations"` // List of locations
	Dates     string   `json:"dates"`     // URL for locations' dates
}

// Dates represents the structure of dates in the Groupie Trackers API
type Dates struct {
	ID    uint     // Dates ID
	Dates []string // List of dates
}

// Relations represents the structure of relations in the Groupie Trackers API
type Relations struct {
	ID             uint                // Relations ID
	DatesLocations map[string][]string // Map of dates and corresponding locations
}

// Data for filtering
type FilterCriteria struct {
	MinCreationDate   string // Date au format ISO 8601 ou autre format adapté
	MaxCreationDate   string
	MinFirstAlbumDate string // Date au format ISO 8601 ou autre format adapté
	MaxFirstAlbumDate string
	Members           []int
	ConcertLocations  string // Liste de lieux séparés par des virgules
}
