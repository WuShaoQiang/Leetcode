package getMinimumDifference

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
	}
)

func TestGetMinimumDifference(t *testing.T) {
	for _, single := range input {
		fmt.Println(getMinimumDifference(single))
	}
}
