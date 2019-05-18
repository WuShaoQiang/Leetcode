package isSameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true

Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false

Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Same Tree.

// 二叉树遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameValue(p, q)
}

func isSameValue(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSameValue(p.Left, q.Left) && isSameValue(p.Right, q.Right)
	}

	return false

}
