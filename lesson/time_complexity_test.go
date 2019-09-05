package lesson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tatoonz/codility/lesson"
)

func TestFrogJmp(t *testing.T) {
	assert.Equal(t, 3, lesson.FrogJmp(10, 85, 30))
	assert.Equal(t, 34, lesson.FrogJmp(5, 105, 3))
}

func TestPermMissingElem(t *testing.T) {
	assert.Equal(t, 4, lesson.PermMissingElem([]int{2, 3, 1, 5}))
	assert.Equal(t, 1, lesson.PermMissingElem([]int{}))
	assert.Equal(t, 2, lesson.PermMissingElem([]int{1}))
}

func TestTapeEquilibrium(t *testing.T) {
	assert.Equal(t, 1, lesson.TapeEquilibrium([]int{3, 1, 2, 4, 3}))
}
