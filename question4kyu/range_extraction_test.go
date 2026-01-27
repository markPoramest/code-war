package question4kyu_test

import (
	"codewar/question4kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeExtraction(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		//assert.Equal(t, "-6,-3-1,5,6", question4kyu.Solution([]int{-6, -3, -2, -1, 0, 1, 5, 6}))
		assert.Equal(t, "-6,-3-1,3-5,7-11,14,15,17-20", question4kyu.Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	})
}
