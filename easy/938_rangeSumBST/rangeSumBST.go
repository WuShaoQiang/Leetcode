package rangeSumBST

/*
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.



Example 1:

Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32

Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23



Note:

    The number of nodes in the tree is at most 10000.
    The final answer is guaranteed to be less than 2^31.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime: 92 ms, faster than 98.06% of Go online submissions for Range Sum of BST.
// Memory Usage: 8.1 MB, less than 70.59% of Go online submissions for Range Sum of BST.

// 这种方法会全部遍历一次
// func rangeSumBST(root *TreeNode, L int, R int) int {
// 	if root == nil {
// 		return 0
// 	}

// 	sum := 0
// 	inOrder(root, L, R, &sum)
// 	return sum
// }

// 这种是按照节点是否为nil来传递
// func inOrder(root *TreeNode, L, R int, sum *int) {
// 	if root.Left != nil {
// 		inOrder(root.Left, L, R, sum)
// 	}

// 	if root.Val >= L && root.Val <= R {

// 		*sum += root.Val
// 	}

// 	if root.Right != nil {
// 		inOrder(root.Right, L, R, sum)
// 	}

// }

// Runtime: 92 ms, faster than 98.06% of Go online submissions for Range Sum of BST.
// Memory Usage: 8.1 MB, less than 56.15% of Go online submissions for Range Sum of BST.

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	inOrder(root, L, R, &sum)
	return sum
}

// 这种是按照节点的值是否符合来传递,这种就不会每次都全部遍历
func inOrder(root *TreeNode, L, R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}

	if root.Val > L {
		inOrder(root.Left, L, R, sum)
	}

	if root.Val < R {
		inOrder(root.Right, L, R, sum)
	}
}
