package solutions

import "github.com/shilin83/leetcode-go/src/structures"

type ListNode = structures.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current, carry := head, uint8(0)

	for l1 != nil || l2 != nil || carry != 0 {
		// * 初始节点的和为进位值
		sum := carry

		// * 如果 l1 链表不为空，将其值加入和，并移动到下一个节点
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// * 如果 l2 链表不为空，将其值加入和，并移动到下一个节点
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// * 创建新节点，值为 sum 除 10 取余
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		// * 计算进位，值为 sum 除 10 取整
		carry = sum / 10
	}

	return head.Next
}
