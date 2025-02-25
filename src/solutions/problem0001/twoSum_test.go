package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [4]struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
	{[]int{1, 2}, 0, []int{}},
}

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run("1.两数之和", func(t *testing.T) {
			actual := twoSum(c.nums, c.target)
			ast.Equal(c.expected, actual)
		})
	}
}
