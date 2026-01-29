package question4kyu_test

import (
	"codewar/question4kyu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	t.Run("should return result correctly", func(t *testing.T) {
		//assert.Equal(t, []string{"ab", "ba"}, question4kyu.Permutations("ab"))
		//assert.Equal(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, question4kyu.Permutations("abc"))
		//assert.Equal(t, []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}, question4kyu.Permutations("aabb"))
		assert.Equal(t, []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda",
			"bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
			question4kyu.Permutations("abcd"))
	})
}
