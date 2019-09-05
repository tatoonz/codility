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
	s := prefixSum(A)

	result := math.MaxFloat64
	for p := 1; p < len(A); p++ {
		diff := math.Abs(float64(s[p-1] - sum(s, p, len(A))))

		if diff < result {
			result = diff
		}
	}

	return int(result)
}

func prefixSum(A []int) []int {
	result := make([]int, len(A))

	total := 0
	for i, n := range A {
		total += n
		result[i] = total
	}

	return result
}

func sum(s []int, from, to int) int {
	return s[to-1] - s[from-1]
}
