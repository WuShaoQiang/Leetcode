package mergeKLists

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime: 252 ms, faster than 21.12% of Go online submissions for Merge k Sorted Lists.
//Memory Usage: 5.6 MB, less than 76.47% of Go online submissions for Merge k Sorted Lists.

//func mergeKLists(lists []*ListNode) *ListNode {
//	last := len(lists) - 1
//	// for case lists == []
//	if last < 0 {
//		return nil
//	}
//	// return one link directly
//	if last == 0 {
//		return lists[0]
//	}
//	res := &ListNode{}
//	head := res
//	insertSort(lists)
//	for lists[last].Val != math.MaxInt64 {
//		res.Next = lists[last]
//		res = res.Next
//		lists[last] = lists[last].Next
//		insertSort(lists)
//	}
//
//	res = nil
//	return head.Next
//}

// descending
func insertSort(lists []*ListNode) {
	// i = 0 for lists[i] == nil
	for i := 0; i < len(lists); i++ {
		picked := lists[i]
		// if nil, we can assign a max value of int64
		if picked == nil {
			picked = &ListNode{Val:math.MaxInt64}
		}
		j := i - 1
		for ; j >= 0 && (picked.Val > lists[j].Val); j-- {
			lists[j+1] = lists[j]
		}
		lists[j+1] = picked
	}
}

//Runtime: 16 ms, faster than 62.16% of Go online submissions for Merge k Sorted Lists.
//Memory Usage: 6 MB, less than 52.94% of Go online submissions for Merge k Sorted Lists.

func mergeKLists(lists []*ListNode) *ListNode{
	n := len(lists)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return lists[0]
	}

	return mergeTwoLinks(mergeKLists(lists[:n/2]),mergeKLists(lists[n/2:]))
}

func mergeTwoLinks(l1,l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		head.Next = mergeTwoLinks(l1.Next,l2)
	} else {
		head = l2
		head.Next = mergeTwoLinks(l1,l2.Next)
	}

	return head
}
