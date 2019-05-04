package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

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
