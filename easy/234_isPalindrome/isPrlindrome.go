package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = reverse(slow.Next)

	for fast != nil {
		if fast.Val != head.Val {
			return false
		}
		fast = fast.Next
		head = head.Next
	}

	return true
}

func reverse(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	if l.Next == nil {
		return l
	}
	var last *ListNode
	curr := l
	for curr.Next != nil {
		tmp := curr.Next
		curr.Next = last
		last = curr
		curr = tmp
	}
	curr.Next = last
	return curr
}

// func isPalindrome(head *ListNode) bool {
// 	v := make([]int, 0)
// 	for head != nil {
// 		v = append(v, head.Val)
// 		head = head.Next
// 	}

// 	l := len(v)

// 	for i := 0; i < l/2; i++ {
// 		if v[i] != v[l-1-i] {
// 			return false
// 		}
// 	}

// 	return true
// }
