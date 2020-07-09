package olc

import (
	olc "github.com/google/open-location-code/go"
)

var digits = [20]string{"2", "3", "4", "5", "6", "7", "8", "9", "C", "F", "G", "H", "J", "M", "P", "Q", "R", "V", "W", "X"}
var Blocks []string
var Subblocks []string

func init() {

	for _, dLat := range digits[1:8] {
		for _, dLon := range digits[1:17] {
			Blocks = append(Blocks, dLat+dLon)
		}
	}

	for _, dLat := range digits {
		for _, dLon := range digits {
			Subblocks = append(Subblocks, dLat+dLon)
		}
	}

}

func CoordinatesToOLC(lat, long float64, length int) string {
	return olc.Encode(lat, long, length)
}
