package utils

import "math"

func RoundAt(rate float64, size int) float64 {
	shift := math.Pow10(size)
	return math.Round(rate*(shift)) / (shift)
}

func CeilAt(rate float64, size int) float64 {
	shift := math.Pow10(size)
	return math.Ceil(rate*shift) / shift
}
