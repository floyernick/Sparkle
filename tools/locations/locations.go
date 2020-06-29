package locations

import (
	olc "github.com/google/open-location-code/go"
)

func CoordinatesToOLC(lat, long float64, length int) string {
	return olc.Encode(lat, long, length)
}
