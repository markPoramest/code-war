package question5kyu_test

import (
	"codewar/question5kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakesAndLadders(t *testing.T) {
	t.Run("should return correctly", func(t *testing.T) {
		var game question5kyu.SnakesLadders

		game.NewGame()

		assert.Equal(t, "Player 1 is on square 38", game.Play(1, 1))
		assert.Equal(t, "Player 1 is on square 44", game.Play(1, 5))
		assert.Equal(t, "Player 2 is on square 31", game.Play(6, 2))
		assert.Equal(t, "Player 1 is on square 25", game.Play(1, 1))
	})
}
