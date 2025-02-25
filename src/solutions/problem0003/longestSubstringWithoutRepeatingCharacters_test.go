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
	ast := assert.New(t)

	for _, c := range cases {
		t.Run("3.无重复字符的最长子串", func(t *testing.T) {
			actual := lengthOfLongestSubstring(c.s)
			ast.Equal(c.expected, actual)
		})
	}
}
