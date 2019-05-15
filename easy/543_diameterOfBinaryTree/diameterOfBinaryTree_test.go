package diameterOfBinaryTree

import (
	"fmt"
	"testing"
)

var (
	input = []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
		&TreeNode{Val: 1},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		&TreeNode{Val: 543, Left: &TreeNode{Val: 384, Right: &TreeNode{Val: 445}}, Right: &TreeNode{Val: 652, Right: &TreeNode{Val: 699}}},
		&TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 13}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: -1}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}}},
	}
)

func TestDiameterOfBinaryTree(t *testing.T) {
	for _, single := range input {
		fmt.Println(diameterOfBinaryTree(single))
	}
}
