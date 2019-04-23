package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	c := head
	for c != nil {
		if c.Next == nil {
			return head
		}
		if c.Val == c.Next.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return head
}
