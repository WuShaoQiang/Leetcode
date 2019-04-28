package maxDepth

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
)

func TestMaxDepth(t *testing.T) {
	fmt.Println(maxDepth(input))
}
