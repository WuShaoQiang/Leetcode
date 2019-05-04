package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val <= q.Val {
		return searchNode(root, p, q)
	}
	return searchNode(root, q, p)
}

//p.Val <= q.Val
func searchNode(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val > root.Val {
		return root
	}

	if root == p || root == q {
		return root
	}

	//都去左边
	if q.Val < root.Val && root.Left != nil {
		return searchNode(root.Left, p, q)
	}

	// 都去右边
	if p.Val > root.Val && root.Right != nil {
		return searchNode(root.Right, p, q)
	}

	return nil
}
