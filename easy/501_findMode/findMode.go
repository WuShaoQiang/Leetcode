package findMode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

// 假定 BST 有如下定义：

//     结点左子树中所含结点的值小于等于当前结点的值
//     结点右子树中所含结点的值大于等于当前结点的值
//     左子树和右子树都是二叉搜索树

// 例如：
// 给定 BST [1,null,2,2],

//    1
//     \
//      2
//     /
//    2

// 返回[2].

// 提示：如果众数超过1个，不需考虑输出顺序

// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

// 执行用时 : 24 ms, 在Find Mode in Binary Search Tree的Go提交中击败了71.43% 的用户
// 内存消耗 : 6.5 MB, 在Find Mode in Binary Search Tree的Go提交中击败了23.53% 的用户

// func findMode(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	m := make(map[int]int)
// 	iteration(root, m)

// 	res := make(map[int][]int)
// 	max := 0
// 	for k, v := range m {
// 		if max <= v {
// 			max = v
// 			res[v] = append(res[v], k)
// 		}
// 	}

// 	return res[max]
// }

// func iteration(root *TreeNode, m map[int]int) {
// 	if root != nil {
// 		m[root.Val]++
// 		iteration(root.Left, m)
// 		iteration(root.Right, m)
// 	}
// }

// Runtime: 8 ms, faster than 100.00% of Go online submissions for Find Mode in Binary Search Tree.
// Memory Usage: 5.8 MB, less than 100.00% of Go online submissions for Find Mode in Binary Search Tree.

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	more := make([]int, 0)

	// 中序遍历的结果存储在res
	orderByMedium(root, &res)

	max := 1
	cur := 1
	for i := 0; i < len(res); i++ {
		// 如果不是最后一个并且前后相同,就继续加
		if i < len(res)-1 && res[i] == res[i+1] {
			cur++
		} else { // 如果是最后一个或者前后不相同,就开始对比
			// 如果是相等,就加进数组
			if cur == max {
				more = append(more, res[i])
			}
			// 如果不相等就重构数组
			if cur > max {
				max = cur
				more = []int{res[i]}
			}
			cur = 1
		}

	}

	return more
}

func orderByMedium(root *TreeNode, res *[]int) {
	if root.Left != nil {
		orderByMedium(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		orderByMedium(root.Right, res)
	}
}
