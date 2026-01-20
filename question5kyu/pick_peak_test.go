package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickPeaks(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.Equal(t, question5kyu.PosPeaks{Pos: []int{3}, Peaks: []int{5}}, question5kyu.PickPeaks([]int{0, 1, 2, 5, 1, 0}))
		assert.Equal(t, question5kyu.PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}, question5kyu.PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))
		assert.Equal(t, question5kyu.PosPeaks{Pos: []int{1}, Peaks: []int{2}}, question5kyu.PickPeaks([]int{1, 2, 2, 2, 1}))
		assert.Equal(t, question5kyu.PosPeaks{Pos: []int{2}, Peaks: []int{3}}, question5kyu.PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2}))
		assert.Equal(t, question5kyu.PosPeaks{Pos: []int{5}, Peaks: []int{10, 2}}, question5kyu.PickPeaks([]int{10, -1, -4, -3, 2, 1, 10}))
	})
}
