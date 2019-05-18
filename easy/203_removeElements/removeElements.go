package removeElements

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5

*/

// Runtime: 8 ms, faster than 98.48% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.6 MB, less than 84.18% of Go online submissions for Remove Linked List Elements.

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	tmp := head

	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return head
}
