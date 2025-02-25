package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [4]struct {
	nums1, nums2 []int
	expected     float64
}{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{1, 2, 3, 4}, []int{5, 6, 7}, 4.0},
	{[]int{0}, []int{0}, 0.0},
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run("4.寻找两个正序数组的中位数", func(t *testing.T) {
			actual := findMedianSortedArrays(c.nums1, c.nums2)
			ast.Equal(c.expected, actual)
		})
	}
}
