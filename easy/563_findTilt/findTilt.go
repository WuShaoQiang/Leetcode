package findTilt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，计算整个树的坡度。

// 一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。

// 整个树的坡度就是其所有节点的坡度之和。

// 示例:

// 输入:
//          1
//        /   \
//       2     3
// 输出: 1
// 解释:
// 结点的坡度 2 : 0
// 结点的坡度 3 : 0
// 结点的坡度 1 : |2-3| = 1
// 树的坡度 : 0 + 0 + 1 = 1

// 注意:

//     任何子树的结点的和不会超过32位整数的范围。
//     坡度的值不会超过32位整数的范围。

// 执行用时 : 8 ms, 在Binary Tree Tilt的Go提交中击败了100.00% 的用户
// 内存消耗 : 6.2 MB, 在Binary Tree Tilt的Go提交中击败了31.58% 的用户

func findTilt(root *TreeNode) int {
	var sum int
	findWeight(root, &sum)
	return sum
}

func findWeight(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	left := findWeight(root.Left, sum)
	right := findWeight(root.Right, sum)

	*sum += abs(left - right)
	return root.Val + left + right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
