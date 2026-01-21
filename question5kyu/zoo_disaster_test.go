package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhoEatsWho(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		assert.Equal(t, []string{
			"fox,bug,chicken,grass,sheep",
			"chicken eats bug",
			"fox eats chicken",
			"sheep eats grass",
			"fox eats sheep",
			"fox"}, question5kyu.WhoEatsWho("fox,bug,chicken,grass,sheep"))

		assert.Equal(t, []string{"chicken,fox,leaves,bug,grass,sheep",
			"fox eats chicken",
			"bug eats leaves",
			"sheep eats grass",
			"fox,bug,sheep"}, question5kyu.WhoEatsWho("chicken,fox,leaves,bug,grass,sheep"))
	})
}
