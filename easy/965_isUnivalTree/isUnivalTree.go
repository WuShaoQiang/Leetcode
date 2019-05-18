package isUnivalTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.



Example 1:

Input: [1,1,1,1,1,null,1]
Output: true

Example 2:

Input: [2,2,2,5,2]
Output: false



Note:

    The number of nodes in the given tree will be in the range [1, 100].
    Each node's value will be an integer in the range [0, 99].


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Univalued Binary Tree.
// Memory Usage: 2.3 MB, less than 63.06% of Go online submissions for Univalued Binary Tree.

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	val := root.Val
	return isSame(root.Left, val) && isSame(root.Right, val)
}

func isSame(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	}

	return isSame(root.Left, val) && isSame(root.Right, val)
}
