package structures

type ListNode struct {
	Val  uint8
	Next *ListNode
}

func ToList(nums []uint8) *ListNode {
	var head *ListNode

	for _, num := range nums {
		node := &ListNode{
			Val:  num,
			Next: head,
		}

		head = node
	}

	return head
}

func ToArray(head *ListNode) []uint8 {
	var nums []uint8

	for nil != head {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}
