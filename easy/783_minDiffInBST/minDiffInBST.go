package minDiffInBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a Binary Search Tree (BST) with the root node root, return the minimum difference between the values of any two different nodes in the tree.

Example :

Input: root = [4,2,6,1,3,null,null]
Output: 1
Explanation:
Note that root is a TreeNode object, not an array.

The given tree [4,2,6,1,3,null,null] is represented by the following diagram:

          4
        /   \
      2      6
     / \
    1   3

while the minimum difference in this tree is 1, it occurs between node 1 and node 2, also between node 3 and node 2.

Note:

    The size of the BST will be between 2 and 100.
    The BST is always valid, each node's value is an integer, and each node's value is different.
*/

// 执行用时 : 4 ms, 在Minimum Distance Between BST Nodes的Go提交中击败了95.24% 的用户
// 内存消耗 : 2.4 MB, 在Minimum Distance Between BST Nodes的Go提交中击败了76.19% 的用户

// 这个题好像前面也有,如果用数组的思想
func minDiffInBST(root *TreeNode) int {
	res := make([]int, 0)
	inOrder(root, &res)
	return findDiffMin(res)
}

func findDiffMin(res []int) int {
	min := math.MaxInt64
	for i := 1; i < len(res); i++ {
		if min > res[i]-res[i-1] {
			min = res[i] - res[i-1]
		}
	}
	return min
}

func inOrder(root *TreeNode, res *[]int) {
	if root.Left != nil {
		inOrder(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		inOrder(root.Right, res)
	}
}
