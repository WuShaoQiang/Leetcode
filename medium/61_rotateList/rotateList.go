package rotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate List.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Rotate List.

func rotateRight(head *ListNode, k int) *ListNode {
	// first, make link as a cycle link
	if head == nil || head.Next == nil{
		return head
	}

	node := head.Next
	// link length
	length := 2
	// find the last node
	for node.Next != nil {
		node = node.Next
		length++
	}
	// last node point to head, implement cycle link
	node.Next = head

	// calc which node is gonna point to null
	var n int
	// k > length means a full round rotate
	k %= length

	n = length-k

	if n == 0 {
		return head
	}

	pointToNullNode := head
	for i:=1;i<n;i++{
		pointToNullNode = pointToNullNode.Next
	}

	res := pointToNullNode.Next
	pointToNullNode.Next = nil
	return res
}
