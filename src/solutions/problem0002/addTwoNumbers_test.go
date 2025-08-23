package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shilin83/leetcode-go/src/structures"
)

type Case struct {
	l1       []uint8
	l2       []uint8
	expected []uint8
}

var ast = assert.New(nil)

func TestAddTwoNumbers(t *testing.T) {
	cases := [3]Case{
		{[]uint8{2, 4, 3}, []uint8{5, 6, 4}, []uint8{7, 0, 8}},
		{[]uint8{0}, []uint8{0}, []uint8{0}},
		{[]uint8{9, 9, 9, 9, 9, 9, 9}, []uint8{9, 9, 9, 9}, []uint8{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, c := range cases {
		t.Run("2.两数相加", func(t *testing.T) {
			expected, actual := c.expected, structures.ToArray(
				addTwoNumbers(structures.ToList(c.l1), structures.ToList(c.l2)),
			)

			ast.Equal(expected, actual)
		})
	}
}
