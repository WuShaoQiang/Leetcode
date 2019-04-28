package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
