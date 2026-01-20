package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphanumeric(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.False(t, question5kyu.Alphanumeric(".*?"))
		assert.True(t, question5kyu.Alphanumeric("PassW0rd"))
		assert.False(t, question5kyu.Alphanumeric("     "))
		assert.False(t, question5kyu.Alphanumeric("&)))((("))
		assert.False(t, question5kyu.Alphanumeric("Hello_world"))
	})
}
