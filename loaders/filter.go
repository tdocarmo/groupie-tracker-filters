package loaders

import (
	"groupie-tracker-filters/constants"
	"strconv"
	"strings"
)

// filterData filters the artist data based on the provided criteria
func filterData(data []constants.Artist, criteria constants.FilterCriteria) []constants.Artist {
	var filteredData []constants.Artist

	for _, item := range data {
		minFirstAlbumCriteria, _ := strconv.Atoi(criteria.MinFirstAlbumDate)
		maxFirstAlbumCriteria, _ := strconv.Atoi(criteria.MaxFirstAlbumDate)
		itemFirstAlbum := strings.Split(item.FirstAlbum, "-")
		itemFirstAlbumYear, _ := strconv.Atoi(itemFirstAlbum[2])

		firstAlbumDateOk := (criteria.MinFirstAlbumDate == "" || itemFirstAlbumYear >= minFirstAlbumCriteria) &&
			(criteria.MaxFirstAlbumDate == "" || itemFirstAlbumYear <= maxFirstAlbumCriteria)

		minCreationDateCriteria, _ := strconv.Atoi(criteria.MinCreationDate)
		maxCreationDateCriteria, _ := strconv.Atoi(criteria.MaxCreationDate)
		creationDateOk := (criteria.MinCreationDate == "" || item.CreationDate >= minCreationDateCriteria) &&
			(criteria.MaxCreationDate == "" || item.CreationDate <= maxCreationDateCriteria)

		// Filtre par nombre de membres
		membersOk := len(criteria.Members) == 0
		for _, member := range criteria.Members {
			if len(item.Members) == member {
				membersOk = true
				break
			}
		}

		// Filtre par localisations
		locationsOk := criteria.ConcertLocations == ""
		// Search for matching Id's between artists and locations.
		for _, location := range locations {
			if location.ID == item.ID {
				for _, loc := range location.Locations {
					locFilter := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(loc, "-", " "), "_", " "))
					criteriaFilter := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(criteria.ConcertLocations, "-", " "), "_", " "))
					if strings.Contains(locFilter, criteriaFilter) {
						locationsOk = true
						break
					}
				}
			}
		}

		if creationDateOk && firstAlbumDateOk && membersOk && locationsOk {
			filteredData = append(filteredData, item)
		}
	}
	return filteredData
}
