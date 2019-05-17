package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.5 MB, less than 90.64% of Go online submissions for Merge Two Sorted Lists.

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l4 := &ListNode{}
	l3 := &ListNode{}
	l4.Next = l3
	for l1 != nil || l2 != nil {
		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}
		if l1.Val <= l2.Val {
			l3.Next = l1
			l3 = l3.Next
			if l1.Next == nil {
				l3.Next = l2
				return l4.Next.Next
			}
			l1 = l1.Next
		} else {
			l3.Next = l2
			l3 = l3.Next
			if l2.Next == nil {
				l3.Next = l1
				return l4.Next.Next
			}
			l2 = l2.Next
		}
	}
	return nil
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.5 MB, less than 90.64% of Go online submissions for Merge Two Sorted Lists.

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	ret := ListNode{}
// 	c := &ret

// 	for l1 != nil && l2 != nil {
// 		if l1.Val <= l2.Val {
// 			c.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			c.Next = l2
// 			l2 = l2.Next
// 		}
// 		c = c.Next
// 	}

// 	if l1 == nil {
// 		c.Next = l2
// 	} else if l2 == nil {
// 		c.Next = l1
// 	}

// 	return ret.Next
// }
