package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

For example,

Given the tree:
        4
       / \
      2   7
     / \
    1   3

And the value to search: 2

You should return this subtree:

      2
     / \
    1   3

In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL.

Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as [], not null.

*/

// Runtime: 28 ms, faster than 66.23% of Go online submissions for Search in a Binary Search Tree.
// Memory Usage: 6.7 MB, less than 5.38% of Go online submissions for Search in a Binary Search Tree.

// 比较简单
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}
