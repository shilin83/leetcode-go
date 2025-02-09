package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	cases := [4]Case{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2}, 4, []int{}},
	}

	ast := assert.New(t)

	for _, c := range cases {
		t.Run("1.两数之和", func(t *testing.T) {
			ast.Equal(c.expected, twoSum(c.nums, c.target))
		})
	}
}
