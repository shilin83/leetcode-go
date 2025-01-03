package solutions

import "github.com/shilin83/leetcode-go/src/structures"

type ListNode = structures.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	current, carry := head, uint8(0)

	// * 遍历两个链表，计算每个节点的和，并与当前进位值相加
	// * 遍历结束后，如果还有进位值，则在链表末尾添加新的节点
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}

	return head.Next
}
