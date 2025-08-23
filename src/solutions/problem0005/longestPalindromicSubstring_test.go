package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	s        string
	expected string
}

var ast = assert.New(nil)

func TestLongestPalindrome(t *testing.T) {
	cases := [2]Case{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, c := range cases {
		t.Run("5.最长回文子串", func(t *testing.T) {
			expected, actual := c.expected, longestPalindrome(c.s)

			ast.Equal(expected, actual)
		})
	}
}
