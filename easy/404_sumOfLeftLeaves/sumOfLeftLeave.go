package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Left Leaves.
// Memory Usage: 2.7 MB, less than 79.63% of Go online submissions for Sum of Left Leaves.

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	search(root, false, &sum)
	return sum

}

func search(root *TreeNode, flag bool, sum *int) {
	if root == nil {
		return
	}

	if flag && root.Left == nil && root.Right == nil {
		*sum += root.Val
	}

	if root.Left != nil {
		search(root.Left, true, sum)
	}

	if root.Right != nil {
		search(root.Right, false, sum)
	}
}
