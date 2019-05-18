package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its depth = 3.

*/

// Runtime: 4 ms, faster than 99.27% of Go online submissions for Maximum Depth of Binary Tree.
// Memory Usage: 4.4 MB, less than 93.21% of Go online submissions for Maximum Depth of Binary Tree.

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//DFS
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	if root.Left != nil && root.Right != nil {
// 		left := maxDepth(root.Left) + 1
// 		right := maxDepth(root.Right) + 1
// 		if left > right {
// 			return left
// 		}

// 		return right
// 	}

// 	if root.Left != nil {
// 		return maxDepth(root.Left) + 1
// 	}

// 	if root.Right != nil {
// 		return maxDepth(root.Right) + 1
// 	}
// 	return 1
// }
