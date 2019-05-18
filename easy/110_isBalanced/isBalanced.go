package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

    a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7

Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4

Return false.
*/

// Runtime: 8 ms, faster than 98.32% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 6.1 MB, less than 93.64% of Go online submissions for Balanced Binary Tree.

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == 0 && right == 0 {
		return true
	}

	// fmt.Printf("left : %v,right : %v\n", left, right)

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
