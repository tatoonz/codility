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

// TapeEquilibrium minimize the value |(A[0] + ... + A[P-1]) - (A[P] + ... + A[N-1])|.
func TapeEquilibrium(A []int) int {
	total := 0
	for i := range A {
		total += A[i]
	}

	result := math.MaxFloat64
	left := 0
	for p := 1; p < len(A); p++ {
		left += A[p-1]
		right := total - left

		diff := math.Abs(float64(left - right))
		if diff < result {
			result = diff
		}
	}

	return int(result)
}
