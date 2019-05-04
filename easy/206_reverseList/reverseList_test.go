package reverseList

import (
	"fmt"
	"testing"
)

var (
	input = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}
)

func TestReverseList(t *testing.T) {
	result := reverseList(input)
	fmt.Println(result.Next.Next.Val)
}
