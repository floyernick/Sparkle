package geohash

import "github.com/mmcloughlin/geohash"

const MAX_LENGTH = 10

func CoordinatesToGeohash(lat, long float64, length int) string {
	return geohash.Encode(lat, long)[:length]
}

func GeohashToCoordinates(hash string) (lat, long float64) {
	return geohash.DecodeCenter(hash)
}

func GetClickableLength(zoom int) int {
	switch {
	case zoom >= 17:
		return 8
	case zoom >= 14:
		return 7
	case zoom >= 11:
		return 6
	case zoom >= 9:
		return 5
	case zoom >= 6:
		return 4
	default:
		return 3
	}
}

func GetDisplayableLength(zoom int) int {
	switch {
	case zoom > 14:
		return 8
	case zoom > 11:
		return 7
	case zoom > 9:
		return 6
	case zoom > 6:
		return 5
	default:
		return 4
	}
}

func GetParentLength(zoom int) int {
	switch {
	case zoom > 14:
		return 5
	case zoom > 11:
		return 4
	case zoom > 9:
		return 3
	case zoom > 6:
		return 2
	default:
		return 0
	}
}
