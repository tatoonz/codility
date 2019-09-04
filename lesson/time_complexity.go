package lesson

import (
	"math"
)

// FrogJmp count minimal number of jumps from position X to Y.
func FrogJmp(X, Y, D int) int {
	result := math.Ceil(float64(Y-X) / float64(D))
	return int(result)
}
