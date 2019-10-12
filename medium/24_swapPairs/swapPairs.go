package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
//Memory Usage: 2.1 MB, less than 33.33% of Go online submissions for Swap Nodes in Pairs.

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextRound := head.Next.Next
	l1 := head
	l2 := head.Next
	l2.Next = l1
	l1.Next = swapPairs(nextRound)

	return l2
}
