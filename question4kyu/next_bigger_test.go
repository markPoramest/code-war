package question4kyu_test

import (
	"codewar/question4kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextBigger(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.Equal(t, 2071, question4kyu.NextBigger(2017))
		assert.Equal(t, 531, question4kyu.NextBigger(513))
		assert.Equal(t, -1, question4kyu.NextBigger(54321))
		assert.Equal(t, 414, question4kyu.NextBigger(144))
		assert.Equal(t, 1234567908, question4kyu.NextBigger(1234567890))
		assert.Equal(t, 59884848483559, question4kyu.NextBigger(59884848459853))
	})
}
