package deleteNode

import (
	"fmt"
	"testing"
)

var (
	l    = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: node}}
	node = &ListNode{Val: 1, Next: &ListNode{Val: 9}}
)

func TestDeleteNode(t *testing.T) {
	deleteNode(node)
	fmt.Println(l.Next.Next.Val)
}
