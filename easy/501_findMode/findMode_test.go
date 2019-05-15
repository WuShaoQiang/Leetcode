package findMode

import (
	"fmt"
	"testing"
)

var (
	input = []*TreeNode{
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}, Left: &TreeNode{Val: 1}},
		&TreeNode{Val: 1},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	}
)

func TestFindMode(t *testing.T) {
	for _, single := range input {

		fmt.Println(findMode(single))
	}
}
