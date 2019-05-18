package levelOrderBottom

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its bottom-up level order traversal as:

[
  [15,7],
  [9,20],
  [3]
]
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
// Memory Usage: 5.9 MB, less than 95.56% of Go online submissions for Binary Tree Level Order Traversal II.

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return nil
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		len := len(que)
		var level []int
		for i := 0; i < len; i++ {
			node := que[0]
			que = que[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		ret = append(ret, level)
	}
	var ans [][]int
	for i := len(ret) - 1; i >= 0; i-- {
		ans = append(ans, ret[i])
	}
	return ans
}
