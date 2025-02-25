package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToList(nums []int) *ListNode {
	var head *ListNode

	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  nums[i],
			Next: head,
		}

		head = node
	}

	return head
}

func ListToArray(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
