package isSameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameValue(p, q)
}

func isSameValue(p *TreeNode, q *TreeNode) bool {
	// if (p == nil && q != nil) || (p != nil && q == nil) {
	// 	return false
	// }

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
