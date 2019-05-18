package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.5 MB, less than 67.14% of Go online submissions for Reverse Linked List.

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// curr := head
	// return reverse(nil, curr)
	return reverse(head)
}

// func reverse(parent, curr *ListNode) *ListNode {
// 	if curr.Next == nil {
// 		curr.Next = parent
// 		return curr
// 	}

// 	result := reverse(curr, curr.Next)
// 	curr.Next = parent
// 	return result
// }

func reverse(head *ListNode) *ListNode {
	var parent *ListNode
	curr := head
	for curr.Next != nil {
		tmp := curr.Next
		curr.Next = parent
		parent = curr
		curr = tmp
	}
	curr.Next = parent
	return curr
}
