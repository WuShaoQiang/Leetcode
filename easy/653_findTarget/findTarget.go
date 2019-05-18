package findTarget

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True



Example 2:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
*/

// Runtime: 24 ms, faster than 97.99% of Go online submissions for Two Sum IV - Input is a BST.
// Memory Usage: 7.2 MB, less than 71.54% of Go online submissions for Two Sum IV - Input is a BST.

// 第一思路,换成数组,利用双边夹方法
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	res := make([]int, 0)
	inOrder(root, &res)

	// 确定首尾索引
	head := 0
	tail := len(res) - 1

	for head < tail {
		sum := res[head] + res[tail]
		if sum == k {
			return true
		}

		if sum > k {
			tail--
		} else {
			head++
		}
	}

	return false
}

// 执行用时 : 40 ms, 在Two Sum IV - Input is a BST的Go提交中击败了98.11% 的用户
// 内存消耗 : 7.3 MB, 在Two Sum IV - Input is a BST的Go提交中击败了16.00% 的用户

func inOrder(root *TreeNode, res *[]int) {
	if root.Left != nil {
		inOrder(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		inOrder(root.Right, res)
	}
}

// 执行用时 : 44 ms, 在Two Sum IV - Input is a BST的Go提交中击败了84.91% 的用户
// 内存消耗 : 7.1 MB, 在Two Sum IV - Input is a BST的Go提交中击败了36.00% 的用户

// 中序遍历
// func inOrder(root *TreeNode) []int {
// 	res := make([]int, 0)

// 	if root.Left != nil {
// 		res = append(res, inOrder(root.Left)...)
// 	}

// 	res = append(res, root.Val)

// 	if root.Right != nil {
// 		res = append(res, inOrder(root.Right)...)
// 	}

// 	return res
// }
