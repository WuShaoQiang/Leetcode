package binaryTreePaths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
// Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Binary Tree Paths.

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
