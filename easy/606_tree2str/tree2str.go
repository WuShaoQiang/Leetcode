package tree2str

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 映射关系的空括号对。

// 示例 1:

// 输入: 二叉树: [1,2,3,4]
//        1
//      /   \
//     2     3
//    /
//   4

// 输出: "1(2(4))(3)"

// 解释: 原本将是“1(2(4)())(3())”，
// 在你省略所有不必要的空括号对之后，
// 它将是“1(2(4))(3)”。

// 示例 2:

// 输入: 二叉树: [1,2,3,null,4]
//        1
//      /   \
//     2     3
//      \
//       4

// 输出: "1(2()(4))(3)"

// 解释: 和第一个示例相似，
// 除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。

// 执行用时 : 40 ms, 在Construct String from Binary Tree的Go提交中击败了18.52% 的用户
// 内存消耗 : 8.8 MB, 在Construct String from Binary Tree的Go提交中击败了26.67% 的用户

// func tree2str(t *TreeNode) string {
//     if t == nil{
//         return ""
//     }
// 	return createStr(t)
// }

// func createStr(t *TreeNode) string {
// 	if t.Left != nil && t.Right != nil {
// 		return fmt.Sprintf("%d(%s)(%s)", t.Val, createStr(t.Left), createStr(t.Right))
// 	}

// 	if t.Left != nil && t.Right == nil {
// 		return fmt.Sprintf("%d(%s)", t.Val, createStr(t.Left))
// 	}

// 	if t.Left == nil && t.Right != nil {
// 		return fmt.Sprintf("%d()(%s)", t.Val, createStr(t.Right))
// 	}

// 	return fmt.Sprintf("%d", t.Val)

// }

// 执行用时 : 16 ms, 在Construct String from Binary Tree的Go提交中击败了100.00% 的用户
// 内存消耗 : 7.1 MB, 在Construct String from Binary Tree的Go提交中击败了93.33% 的用户

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left != nil && t.Right != nil {
		// return fmt.Sprintf("%d(%s)(%s)", t.Val, createStr(t.Left), createStr(t.Right))
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
	}

	if t.Left != nil && t.Right == nil {
		// return fmt.Sprintf("%d(%s)", t.Val, createStr(t.Left))
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}

	if t.Right != nil && t.Left == nil {
		// return fmt.Sprintf("%d(%s)", t.Val, createStr(t.Left))
		return strconv.Itoa(t.Val) + "()(" + tree2str(t.Right) + ")"
	}

	return strconv.Itoa(t.Val)
}
