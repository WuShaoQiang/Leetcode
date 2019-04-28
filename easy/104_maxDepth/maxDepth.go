package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
