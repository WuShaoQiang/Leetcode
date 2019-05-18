package leafSimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Leaf-Similar Trees.
// Memory Usage: 2.5 MB, less than 96.15% of Go online submissions for Leaf-Similar Trees.

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	res1 := make([]int, 0)
	res2 := make([]int, 0)

	findLeaf(root1, &res1)
	findLeaf(root2, &res2)

	if len(res1) != len(res2) {
		return false
	}

	for i := 0; i < len(res1); i++ {
		if res1[i] != res2[i] {
			return false
		}
	}

	return true
}

func findLeaf(root *TreeNode, res *[]int) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
	}

	if root.Left != nil {
		findLeaf(root.Left, res)
	}

	if root.Right != nil {
		findLeaf(root.Right, res)
	}
}
