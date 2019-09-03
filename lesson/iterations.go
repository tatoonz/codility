package lesson

import "strconv"

// BinaryGap finds longest sequence of zeros in binary representation of an integer.
func BinaryGap(N int) int {
	largestGap := 0
	gap := 0

	for _, c := range strconv.FormatInt(int64(N), 2) {
		if c != '1' {
			gap++
			continue
		}

		if gap > largestGap {
			largestGap = gap
		}

		gap = 0
	}

	return largestGap
}
