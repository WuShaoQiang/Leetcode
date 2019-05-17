package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2

Example 2:

Input: 1->1->2->3->3
Output: 1->2->3

*/

// Runtime: 4 ms, faster than 97.90% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.1 MB, less than 42.77% of Go online submissions for Remove Duplicates from Sorted List.

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	c := head
	for c.Next != nil {
		if c.Val == c.Next.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return head
}

// Runtime: 4 ms, faster than 97.90% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 12.14% of Go online submissions for Remove Duplicates from Sorted List.

// 递归做法
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	if head.Next == nil {
// 		return head
// 	}

// 	head.Next = deleteDuplicates(head.Next)

// 	if head.Val == head.Next.Val {
// 		head.Next = head.Next.Next
// 	}

// 	return head
// }
