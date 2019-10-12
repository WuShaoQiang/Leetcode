package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
//Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Remove Nth Node From End of List.

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for i:=0;i<n;i++{
		fast = fast.Next
	}
	// if fast is nil means the link length is n, so remove the first node
	if fast == nil {
		return head.Next
	}

	// now the number of nodes is n between fast and slow
	fast = fast.Next

	// move until fast is reach the nil
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove the next node of slow
	slow.Next = slow.Next.Next

	return head
}
