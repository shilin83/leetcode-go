package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [3]struct {
	s        string
	expected string
}{
	{"babad", "aba"},
	{"cbbd", "bb"},
	{"a", "a"},
}

func TestLongestPalindromicSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run("5.最长回文子串", func(t *testing.T) {
			actual := longestPalindrome(c.s)
			ast.Equal(c.expected, actual)
		})
	}
}
