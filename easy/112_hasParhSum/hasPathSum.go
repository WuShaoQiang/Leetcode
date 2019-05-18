package hasParhSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1

return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Path Sum.
// Memory Usage: 4.8 MB, less than 59.74% of Go online submissions for Path Sum.

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return count(root, sum)
}

func count(root *TreeNode, sum int) bool {
	if root != nil {
		sum = sum - root.Val
		// 右边为空，左边有值并不是叶子节点
		if root.Right == nil && root.Left != nil {
			return count(root.Left, sum)
		}
		// 左边为空，右边有值也不是叶子节点
		if root.Right != nil && root.Left == nil {
			return count(root.Right, sum)
		}
		return count(root.Left, sum) || count(root.Right, sum)
	}
	//必须是叶子节点
	if root == nil {
		if sum == 0 {
			return true
		}
	}
	return false
}
