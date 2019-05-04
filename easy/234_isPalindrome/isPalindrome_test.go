package isPalindrome

import (
	"fmt"
	"testing"
)

var (
	input = []*ListNode{
		// &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
	}
)

func TestIsPalindrome(t *testing.T) {
	for _, single := range input {
		fmt.Println(isPalindrome(single))
	}
}
