package binaryTreePaths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]string, 0)

	findPaths(root, "", &res)
	return res
}

func findPaths(root *TreeNode, str string, result *[]string) {
	if root.Right == nil && root.Left == nil {
		*result = append(*result, str+strconv.Itoa(root.Val))
	}

	if root.Left != nil {
		findPaths(root.Left, str+strconv.Itoa(root.Val)+"->", result)
	}

	if root.Right != nil {
		findPaths(root.Right, str+strconv.Itoa(root.Val)+"->", result)
	}
}
