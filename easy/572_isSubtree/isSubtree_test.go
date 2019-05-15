package isSubtree

import (
	"fmt"
	"testing"
)

var (
	input = []*TreeNode{
		&TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}},
		&TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		&TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 5}},
		&TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		&TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}},
		&TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
	}
)

func TestIsSubtree(t *testing.T) {
	for i := 0; i < len(input); i += 2 {
		fmt.Println(isSubtree(input[i], input[i+1]))
	}
}
