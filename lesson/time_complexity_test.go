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
