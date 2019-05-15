package getMinimumDifference

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

// 示例 :

// 输入:

//    1
//     \
//      3
//     /
//    2

// 输出:
// 1

// 解释:
// 最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

// 注意: 树中至少有2个节点。

// 执行用时 : 20 ms, 在Minimum Absolute Difference in BST的Go提交中击败了100.00% 的用户
// 内存消耗 : 6.4 MB, 在Minimum Absolute Difference in BST的Go提交中击败了85.71% 的用户

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := make([]int, 0)
	orderByIn(root, &res)

	fmt.Println(res)

	min := math.MaxInt64

	for i := len(res) - 1; i > 0; i-- {
		if min > res[i]-res[i-1] {
			min = res[i] - res[i-1]
		}
	}

	return min

}

func orderByIn(root *TreeNode, res *[]int) {
	if root.Left != nil {
		orderByIn(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		orderByIn(root.Right, res)
	}
}
