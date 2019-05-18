package invertTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9

Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1

Trivia:
This problem was inspired by this original tweet by Max Howell:

    Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.2 MB, less than 64.58% of Go online submissions for Invert Binary Tree.

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
