package tests

import (
	"testing"

	"github.com/shilin83/leetcode-go/src/solutions"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{}, 0, []int{}},
	}

	ast := assert.New(t)

	for _, c := range cases {
		t.Run("1.两数之和", func(t *testing.T) {
			actual := solutions.TwoSum(c.nums, c.target)

			ast.Equal(c.expected, actual)
		})
	}
}
