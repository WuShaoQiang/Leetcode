package hasCycle

import (
	"fmt"
	"testing"
)

var (
	input = []*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
	}
)

func TestHasCycle(t *testing.T) {
	fmt.Println(hasCycle(input[0]))
}
