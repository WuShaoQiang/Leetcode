package isBalanced

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == 0 && right == 0 {
		return true
	}

	fmt.Printf("left : %v,right : %v\n", left, right)

	dis := abs(left - right)

	if dis > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)

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

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// func isBalancedNode(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}

// 	//其中之一是nil
// 	if root.Left == nil {
// 		if !(root.Right.Left == nil && root.Right.Right == nil) {
// 			return false
// 		}
// 	}

// 	if root.Right == nil {
// 		if !(root.Left.Left == nil && root.Left.Right == nil) {
// 			return false
// 		}
// 	}

// 	return isBalancedNode(root.Left) && isBalancedNode(root.Right)
// }
