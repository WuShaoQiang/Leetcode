package sumRootToLeaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.

Return the sum of these numbers.



Example 1:

Input: [1,0,1,0,1,0,1]
Output: 22
Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22



Note:

    The number of nodes in the tree is between 1 and 1000.
    node.val is 0 or 1.
    The answer will not exceed 2^31 - 1.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Root To Leaf Binary Numbers.
// Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Sum of Root To Leaf Binary Numbers.

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	toInt(root, 0, &sum)

	return sum
}

func toInt(root *TreeNode, num int, res *int) {
	// 叶子
	if root.Left == nil && root.Right == nil {
		num = num | root.Val
		*res += num
		return
	}

	// 非叶子
	num = num | root.Val

	if root.Left != nil {
		toInt(root.Left, num<<1, res)
	}

	if root.Right != nil {
		toInt(root.Right, num<<1, res)
	}
}
