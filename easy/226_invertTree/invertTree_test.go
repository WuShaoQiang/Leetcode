package invertTree

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
)

func TestInvertTree(t *testing.T) {
	result := invertTree(input)
	fmt.Println(result.Right.Right.Val)
}
