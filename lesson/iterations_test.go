package lesson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/codility/lesson"
)

func TestBinaryGap(t *testing.T) {
	t.Run("FirstExample", func(t *testing.T) {
		actual := lesson.BinaryGap(1041)
		expected := 5
		assert.Equal(t, expected, actual)
	})

	t.Run("SecondExample", func(t *testing.T) {
		actual := lesson.BinaryGap(32)
		expected := 0
		assert.Equal(t, expected, actual)
	})
}
