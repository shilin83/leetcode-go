package structures

import "slices"

type ListNode struct {
	Val  uint8
	Next *ListNode
}

func Int2List(nums []uint8) *ListNode {
	var head, node *ListNode

	slices.Reverse(nums)

	for _, num := range nums {
		node = &ListNode{
			Val:  num,
			Next: head,
		}

		head = node
	}

	return node
}
