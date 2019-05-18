package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]



Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.



Note:

    All of the nodes' values will be unique.
    p and q are different and both values will exist in the BST.

*/

// Runtime: 24 ms, faster than 73.95% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
// Memory Usage: 6.8 MB, less than 82.30% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.

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
