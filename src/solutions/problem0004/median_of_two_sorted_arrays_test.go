package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [2]struct {
	nums1    []int
	nums2    []int
	expected float64
}{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	for _, c := range cases {
		t.Run("4.寻找两个正序数组的中位数", func(t *testing.T) {
			actual := findMedianSortedArrays(c.nums1, c.nums2)

			assert.Equal(t, c.expected, actual)
		})
	}
}
