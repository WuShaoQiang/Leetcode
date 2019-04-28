package isSameTree

import (
	"fmt"
	"testing"
)

var (
	p = []*TreeNode{
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
	}

	q = []*TreeNode{
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	}
)

func TestIsSameTree(t *testing.T) {
	for i := 0; i < len(p); i++ {
		fmt.Println(isSameTree(p[i], q[i]))
	}
}
