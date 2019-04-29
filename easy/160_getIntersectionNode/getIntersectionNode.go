package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA := headA
	tmpB := headB
	dis := 0

	for tmpA != nil && tmpB != nil {
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}
	for tmpB != nil {
		tmpB = tmpB.Next
		dis--
	}

	for tmpA != nil {
		tmpA = tmpA.Next
		dis++
	}

	tmpA = headA
	tmpB = headB

	if dis > 0 {
		for dis > 0 {
			tmpA = tmpA.Next
			dis--
		}
	}

	if dis < 0 {
		for dis < 0 {
			tmpB = tmpB.Next
			dis++
		}
	}

	for tmpA != nil && tmpB != nil {
		if tmpA == tmpB {
			return tmpA
		}
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return nil
}
