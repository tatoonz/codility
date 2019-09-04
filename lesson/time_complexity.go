package lesson

import (
	"math"
)

// FrogJmp count minimal number of jumps from position X to Y.
func FrogJmp(X, Y, D int) int {
	result := math.Ceil(float64(Y-X) / float64(D))
	return int(result)
}

// PermMissingElem finds the missing element in a given permutation.
func PermMissingElem(A []int) int {
	n := len(A) + 1

	expectedTotal := n * (n + 1) / 2

	total := 0
	for i := range A {
		total += A[i]
	}

	return expectedTotal - total
}
