package tests

import (
	"testing"

	"github.com/shilin83/leetcode-go/src/solutions"
	"github.com/shilin83/leetcode-go/src/structures"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := [3]struct {
		nums1, nums2, expected []uint8
	}{
		{[]uint8{2, 4, 3}, []uint8{5, 6, 4}, []uint8{7, 0, 8}},
		{[]uint8{0}, []uint8{0}, []uint8{0}},
		{[]uint8{9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9}, []uint8{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	ast := assert.New(t)

	for _, c := range cases {
		t.Run("2.两数相加", func(t *testing.T) {
			actual := solutions.AddTwoNumbers(structures.Int2List(c.nums1), structures.Int2List(c.nums2))

			ast.Equal(structures.Int2List(c.expected), actual)
		})
	}
}
