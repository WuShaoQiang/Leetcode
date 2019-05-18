package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，它的每个结点都存放着一个整数值。

// 找出路径和等于给定数值的路径总数。

// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

// 示例：

// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

//       10
//      /  \
//     5   -3
//    / \    \
//   3   2   11
//  / \   \
// 3  -2   1

// 返回 3。和等于 8 的路径有:

// 1.  5 -> 3
// 2.  5 -> 2 -> 1
// 3.  -3 -> 11

// Runtime: 12 ms, faster than 85.54% of Go online submissions for Path Sum III.
// Memory Usage: 4.3 MB, less than 89.92% of Go online submissions for Path Sum III.

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	iteration(root, sum, &count)
	return count
}

func iteration(root *TreeNode, target int, count *int) {
	// 每一个不为nil的节点都作为开始节点
	if root != nil {
		findPath(root, 0, target, count)
		iteration(root.Left, target, count)
		iteration(root.Right, target, count)
	}
}

func findPath(root *TreeNode, sum, target int, count *int) {
	//1. 加完后判断
	sum += root.Val
	if sum == target {
		*count++
	}
	//2. 如果判断是count++

	//3. 继续传递

	// 没有孩子了
	if root.Left == nil && root.Right == nil {
		return
	}
	// 有左孩子
	if root.Left != nil {
		findPath(root.Left, sum, target, count)
	}

	// 有右孩子
	if root.Right != nil {
		findPath(root.Right, sum, target, count)
	}
}

// 执行用时 : 12 ms, 在Path Sum III的Go提交中击败了100.00% 的用户
// 内存消耗 : 5.1 MB, 在Path Sum III的Go提交中击败了14.71% 的用户

// 这种方法利用map去记录符合条件的数字的个数,利用map的值传递

// func pathSum(root *TreeNode, sum int) int {
// 	count := 0
// 	m := make(map[int]int)
// 	m[0] = 1
// 	helper(root, 0, sum, m, &count)
// 	return count
// }

// func helper(n *TreeNode, t, s int, m map[int]int, c *int) {
// 	if n == nil {
// 		return
// 	}
// 	t += n.Val
//     if v, _ := m[t-s]; v != 0 {
// 		*c += v
// 	}
// 	m[t] += 1
// 	helper(n.Left, t, s, m, c)
// 	helper(n.Right, t, s, m, c)
//     m[t] -= 1
// }
