package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		sum:= l1.Val + l2.Val
		if carry {
			sum += 1
		}
		if sum > 9 {
			sum -= 10
			carry = true
		} else {
			carry = false
		}
		res.Next = &ListNode{Val:sum}
		l1 = l1.Next
		l2 = l2.Next
		res = res.Next
	}

	for l1 != nil {
		sum := l1.Val
		if carry {
			sum++
		}
		if sum > 9 {
			sum -= 10
			carry = true
		} else {
			carry  =false
		}
		res.Next = &ListNode{Val:sum}
		l1 = l1.Next
		res = res.Next
	}

	for l2 != nil {
		sum := l2.Val
		if carry {
			sum++
		}
		if sum > 9 {
			sum -= 10
			carry = true
		} else {
			carry  =false
		}
		res.Next = &ListNode{Val:sum}
		l2 = l2.Next
		res = res.Next
	}

	if carry {
		res.Next = &ListNode{Val:1}
	}

	return head.Next
}

// Other method

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	og := l1
//
//	var saveSpace bool
//	for l1 != nil || l2 != nil {
//		var carry int
//
//		l1.Val += l2.Val
//		if l1.Val >= 10 {
//			l1.Val = l1.Val % 10
//			carry = 1
//		}
//
//		if l1.Next == nil {
//			if l2.Next == nil {
//				if carry == 1 {
//					l2.Val = 1
//					l1.Next = l2
//				}
//				return og
//			}
//
//			l1.Next = l2.Next
//			l2.Next = nil
//		}
//
//		// If l2.Next is nil, we set it to be the carry
//		if l2.Next == nil {
//			l2.Val = carry
//			saveSpace = true
//			carry = 0
//		}
//
//		// If we somehow need to carry here, adjust for that.
//		if carry == 1 {
//			l1.Next.Val += 1
//		}
//
//		l1 = l1.Next
//		if !saveSpace {
//			l2 = l2.Next
//		} else {
//			saveSpace = false
//		}
//
//	}
//
//	return og
//}
