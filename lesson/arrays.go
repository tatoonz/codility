package lesson

// OddOccurrencesInArray finds value that occurs in odd number of elements.
func OddOccurrencesInArray(A []int) int {
	result := 0
	for _, n := range A {
		result ^= n
	}

	return result
}

// CyclicRotation rotates an array to the right by a given number of steps.
func CyclicRotation(A []int, K int) []int {
	if len(A) == K || len(A) <= 1 {
		return A
	}

	result := make([]int, len(A))
	for i, n := range A {
		newI := i + K
		if newI >= len(A) {
			newI -= (newI / len(A)) * len(A)
		}

		result[newI] = n
	}

	return result
}
