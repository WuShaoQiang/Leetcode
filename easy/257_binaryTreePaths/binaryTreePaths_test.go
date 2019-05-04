package binaryTreePaths

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
)

func TestBinaryTreePaths(t *testing.T) {
	fmt.Println(binaryTreePaths(input))
}
