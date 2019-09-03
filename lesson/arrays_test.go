package lesson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/codility/lesson"
)

func TestOddOccurrencesInArray(t *testing.T) {
	tests := map[int][]int{
		7:  []int{9, 3, 9, 3, 9, 7, 9},
		10: []int{10},
	}

	for expected, A := range tests {
		actual := lesson.OddOccurrencesInArray(A)
		assert.Equal(t, expected, actual)
	}
}

func TestCyclicRotation(t *testing.T) {
	t.Run("FirstExample", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{3, 8, 9, 7, 6}, 3)
		expected := []int{9, 7, 6, 3, 8}
		assert.Equal(t, expected, actual)
	})

	t.Run("SecondExample", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{0, 0, 0}, 1)
		expected := []int{0, 0, 0}
		assert.Equal(t, expected, actual)
	})

	t.Run("ThirdExample", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{1, 2, 3, 4}, 4)
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, actual)
	})

	t.Run("EmptyArray", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{}, 10)
		expected := []int{}
		assert.Equal(t, expected, actual)
	})

	t.Run("OneElement", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{3}, 5)
		expected := []int{3}
		assert.Equal(t, expected, actual)
	})

	t.Run("K > len(A)", func(t *testing.T) {
		actual := lesson.CyclicRotation([]int{3, 8, 9, 7, 6}, 7)
		expected := []int{7, 6, 3, 8, 9}
		assert.Equal(t, expected, actual)
	})
}
