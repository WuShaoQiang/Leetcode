package tree2str

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
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
	}
)

func TestTree2Str(t *testing.T) {
	for _, single := range input {
		fmt.Println(tree2str(single))
	}
}
