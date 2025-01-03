package structures

type ListNode struct {
	Val  uint8
	Next *ListNode
}

func Int2List(nums []uint8) *ListNode {
	length := len(nums)

	var head *ListNode

	if length != 0 {
		for i := length - 1; i >= 0; i-- {
			node := &ListNode{Val: nums[i], Next: head}
			head = node
		}
	}

	return head
}
