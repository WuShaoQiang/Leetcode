package findSecondMinimumValue

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes.

Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

If no such second minimum value exists, output -1 instead.

Example 1:

Input:
    2
   / \
  2   5
     / \
    5   7

Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.



Example 2:

Input:
    2
   / \
  2   2

Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.
*/

// 执行用时 : 0 ms, 在Second Minimum Node In a Binary Tree的Go提交中击败了100.00% 的用户
// 内存消耗 : 2 MB, 在Second Minimum Node In a Binary Tree的Go提交中击败了37.50% 的用户

// 有点像层级遍历
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	// 根节点一定是最小值
	min := root.Val
	// 需要找出第二小的值
	secondMin := math.MaxInt64
	queue := []*TreeNode{root}
	res := make([]int, 0)
	for {
		l := len(queue)
		count := l
		for l > 0 {
			// 取出节点
			tmp := queue[l-1]
			// 如果有节点的值大于最小值,那么它就能称为第二小值候选
			if tmp.Val > min {
				res = append(res, tmp.Val)
				// 如果还是与最小值相等,则需要再找下一层,注意nil就不要加进去了
			} else if tmp.Left != nil {
				queue = append(queue, tmp.Left, tmp.Right)
			}
			l--
		}
		queue = queue[count:]
		if len(queue) == 0 {
			break
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] < secondMin {
			secondMin = res[i]
		}
	}
	if secondMin == math.MaxInt64 {
		return -1
	}
	return secondMin
}
