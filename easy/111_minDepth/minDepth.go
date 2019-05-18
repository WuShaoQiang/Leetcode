package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its minimum depth = 2.
*/

// Runtime: 4 ms, faster than 99.64% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 5.2 MB, less than 56.28% of Go online submissions for Minimum Depth of Binary Tree.

func minDepth(root *TreeNode) int {
	return findMinDepth(root)
}

func findMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + min(findMinDepth(root.Left), findMinDepth(root.Right))
}

func min(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a > b {
		return b
	}

	return a
}
