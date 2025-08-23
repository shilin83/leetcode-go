package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	s        string
	expected int
}

var ast = assert.New(nil)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := [3]Case{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, c := range cases {
		t.Run("3.无重复字符的最长子串", func(t *testing.T) {
			expected, actual := c.expected, lengthOfLongestSubstring(c.s)

			ast.Equal(expected, actual)
		})
	}
}
