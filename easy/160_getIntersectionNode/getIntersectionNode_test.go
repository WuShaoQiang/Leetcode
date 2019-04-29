package getIntersectionNode

import (
	"fmt"
	"testing"
)

var (
	intersectionNode = &ListNode{Val: 8, Next: &ListNode{Val: 4}}
	headA            = &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectionNode}}
	headB            = &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: intersectionNode}}}
)

func TestGetIntersectionNode(t *testing.T) {
	result := getIntersectionNode(headA, headB)
	fmt.Println(result.Val)
}
