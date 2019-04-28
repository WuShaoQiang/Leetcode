package hasParhSum

import (
	"fmt"
	"testing"
)

var (
	input  = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	input1 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	input2 = &TreeNode{Val: 1}
	input3 = &TreeNode{Val: 1, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: 3}}
	input4 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}}}
)

func TestHasPathSum(t *testing.T) {
	fmt.Println(hasPathSum(input, 10))
	fmt.Println(hasPathSum(input1, 1))
	fmt.Println(hasPathSum(input1, 2))
	fmt.Println(hasPathSum(input2, 1))
	fmt.Println(hasPathSum(input3, -1))
	fmt.Println(hasPathSum(input4, 6))
}
