package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [3]struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		t.Run("1.两数之和", func(t *testing.T) {
			actual := twoSum(c.nums, c.target)

			assert.Equal(t, c.expected, actual)
		})
	}
}
