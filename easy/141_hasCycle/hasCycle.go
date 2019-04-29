package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

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
