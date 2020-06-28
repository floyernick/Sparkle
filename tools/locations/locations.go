package locations

import (
	olc "github.com/google/open-location-code/go"
)

func ConvertToOLC(lat, long float64) string {
	return olc.Encode(lat, long, 10)
}
