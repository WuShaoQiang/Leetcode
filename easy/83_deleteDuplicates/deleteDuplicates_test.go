package deleteDuplicates

import (
	"fmt"
	"testing"
)

var (
	input = []*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}},
	}
)

func TestDeleteDuplicates(t *testing.T) {
	for _, single := range input {
		result := deleteDuplicates(single)
		for result != nil {
			fmt.Printf("%v-->", result.Val)
			result = result.Next
		}
		fmt.Println("")
	}
}
