package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeating(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.Equal(t, "e", question5kyu.FirstNonRepeating("moonmen"))
		assert.Equal(t, "T", question5kyu.FirstNonRepeating("sTreSS"))
		assert.Equal(t, "#", question5kyu.FirstNonRepeating("~><#~><"))
		assert.Equal(t, "", question5kyu.FirstNonRepeating("aa"))
	})
}
