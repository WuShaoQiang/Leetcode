package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.



Follow up:

Can you solve it using O(1) (i.e. constant) memory?
*/

// Runtime: 4 ms, faster than 99.74% of Go online submissions for Linked List Cycle.
// Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Linked List Cycle.

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow.Val == fast.Val {
			return true
		}
	}
	return false
}

// func hasCycle(head *ListNode) bool {
// 	quick := head
// 	slow := head
// 	for quick != nil && quick.Next != nil {
// 		quick = quick.Next.Next
// 		if slow == quick {
// 			return true
// 		}
// 		slow = slow.Next
// 	}
// 	return false
// }

// func hasCycle(head *ListNode) bool {
// 	hashMap := make(map[*ListNode]bool)

// 	for head != nil {
// 		if _, ok := hashMap[head]; ok {
// 			return true
// 		}

// 		hashMap[head] = true
// 		head = head.Next
// 	}
// 	return false
// }
