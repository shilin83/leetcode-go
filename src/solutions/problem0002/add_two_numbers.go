package problem0002

import "github.com/shilin83/leetcode-go/src/structures"

type ListNode = structures.ListNode

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// * 创建虚拟头节点，作为结果链表的起点
	head := &ListNode{}

	current, carry := head, uint8(0)

	for nil != l1 || nil != l2 || 0 < carry {
		// * 计算两个链表当前节点值与进位值的和，并移动到下一个节点
		sum := carry

		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}

		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}

		// * 创建新节点，并添加到结果链表
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		// * 更新进位值
		carry = sum / 10
	}

	return head.Next
}
