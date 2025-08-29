package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [3]struct {
	s        string
	expected int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, c := range cases {
		t.Run("3.无重复字符的最长子串", func(t *testing.T) {
			actual := lengthOfLongestSubstring(c.s)

			assert.Equal(t, c.expected, actual)
		})
	}
}
