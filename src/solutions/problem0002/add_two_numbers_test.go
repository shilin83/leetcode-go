package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shilin83/leetcode-go/src/structures"
)

var cases = [3]struct {
	l1       []uint8
	l2       []uint8
	expected []uint8
}{
	{[]uint8{2, 4, 3}, []uint8{5, 6, 4}, []uint8{7, 0, 8}},
	{[]uint8{0}, []uint8{0}, []uint8{0}},
	{[]uint8{9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9}, []uint8{8, 9, 9, 9, 0, 0, 0, 1}},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, c := range cases {
		t.Run("2.两数相加", func(t *testing.T) {
			actual := structures.ToArray(
				addTwoNumbers(structures.ToList(c.l1), structures.ToList(c.l2)),
			)

			assert.Equal(t, c.expected, actual)
		})
	}
}
