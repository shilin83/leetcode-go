package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums1    []int
	nums2    []int
	expected float64
}

var ast = assert.New(nil)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := [2]Case{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, c := range cases {
		t.Run("4.寻找两个正序数组的中位数", func(t *testing.T) {
			expected, actual := c.expected, findMedianSortedArrays(c.nums1, c.nums2)

			ast.Equal(expected, actual)
		})
	}
}
