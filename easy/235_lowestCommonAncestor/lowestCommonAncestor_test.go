package lowestCommonAncestor

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 6, Left: p, Right: q}
	p     = &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}
	q     = &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}
)

func TestLowestCommonAncestor(t *testing.T) {
	result := lowestCommonAncestor(input, p, q)
	fmt.Println(result.Val)
}
