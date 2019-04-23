package mergeTwoLists

import (
	"fmt"
	"testing"
)

var (
	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
)

func TestMergeTwoLists(t *testing.T) {
	mergedList := mergeTwoLists(l1, l2)
	for mergedList != nil {
		fmt.Println(mergedList.Val)
		mergedList = mergedList.Next
	}
}
