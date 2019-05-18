package isSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

// 示例 1:
// 给定的树 s:

//      3
//     / \
//    4   5
//   / \
//  1   2

// 给定的树 t：

//    4
//   / \
//  1   2

// 返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

// 示例 2:
// 给定的树 s：

//      3
//     / \
//    4   5
//   / \
//  1   2
//     /
//    0

// 给定的树 t：

//    4
//   / \
//  1   2

// 返回 false。

// 执行用时 : 40 ms, 在Subtree of Another Tree的Go提交中击败了39.62% 的用户
// 内存消耗 : 11.2 MB, 在Subtree of Another Tree的Go提交中击败了5.56% 的用户

// func isSubtree(s *TreeNode, t *TreeNode) bool {

// 	return iteration(s, t)
// }

// func iteration(s *TreeNode, t *TreeNode) bool {
// 	if s == nil {
// 		return false
// 	}

// 	if s.Val == t.Val {
// 		sRes := make([]int, 0)
// 		tRes := make([]int, 0)
// 		orderByIn(s, &sRes)
// 		orderByIn(t, &tRes)
// 		if isSame(sRes, tRes) {
// 			return true
// 		}
// 		return iteration(s.Left, t) || iteration(s.Right, t)
// 	}

// 	return iteration(s.Left, t) || iteration(s.Right, t)

// }

// func orderByIn(root *TreeNode, res *[]int) {
// 	if root.Left != nil {
// 		orderByIn(root.Left, res)
// 	}

// 	*res = append(*res, root.Val)

// 	if root.Right != nil {
// 		orderByIn(root.Right, res)
// 	}
// }

// func isSame(a, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// Runtime: 20 ms, faster than 83.13% of Go online submissions for Subtree of Another Tree.
// Memory Usage: 6.2 MB, less than 62.12% of Go online submissions for Subtree of Another Tree.

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
