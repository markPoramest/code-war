package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}, question5kyu.MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
		assert.Equal(t, []int{0, 0, 0, 0}, question5kyu.MoveZeros([]int{0, 0, 0, 0}))
		assert.Equal(t, []int{}, question5kyu.MoveZeros([]int{}))
	})
}
