package entities

import "Sparkle/tools/geohash"

type Location struct {
	Code        string
	Longitude   float64
	Latitude    float64
	PostsNumber int
}

func LocationsWithCoordinates(locations []Location) []Location {
	newLocations := make([]Location, 0, len(locations))
	for _, location := range locations {
		location.Latitude, location.Longitude = geohash.GeohashToCoordinates(location.Code)
		newLocations = append(newLocations, location)
	}
	return newLocations
}
