package pathSum

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}
)

func TestPathSum(t *testing.T) {
	fmt.Println(pathSum(input, 6))
}
