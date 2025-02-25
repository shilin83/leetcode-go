package problem0002

import (
	"testing"

	"github.com/shilin83/leetcode-go/src/structures"

	"github.com/stretchr/testify/assert"
)

var cases = [3]struct {
	l1, l2, expected []int
}{
	{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	{[]int{0}, []int{0}, []int{0}},
	{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
}

func TestAddTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run("2.两数相加", func(t *testing.T) {
			actual := addTwoNumbers(structures.ArrayToList(c.l1), structures.ArrayToList(c.l2))
			ast.Equal(c.expected, structures.ListToArray(actual))
		})
	}
}
