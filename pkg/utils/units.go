package utils

import "math"

func GibiBytesToBytes(GiB uint8) int64 {
	gibibytes := float64(GiB)
	bytes := int64(gibibytes / math.Pow(1024, 3))
	return bytes
}
