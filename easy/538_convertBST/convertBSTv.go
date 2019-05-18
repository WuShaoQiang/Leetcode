package convertBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

// 例如：

// 输入: 二叉搜索树:
//               5
//             /   \
//            2     13

// 输出: 转换为累加树:
//              18
//             /   \
//           20     13

// 执行用时 : 540 ms, 在Convert BST to Greater Tree的Go提交中击败了35.56% 的用户
// 内存消耗 : 41.9 MB, 在Convert BST to Greater Tree的Go提交中击败了100.00% 的用户

// func convertBST(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	orderByIn(root, nil)

// 	return root
// }

// func orderByIn(root *TreeNode, lastNode *TreeNode) {

// 	if root == nil {
// 		return
// 	}

// 	if lastNode != nil {
// 		root.Val += lastNode.Val + subtreeSum(root.Right)
// 	} else {
// 		root.Val += subtreeSum(root.Right)
// 	}

// 	if root.Right != nil {
// 		orderByIn(root.Right, lastNode)
// 	}

// 	if root.Left != nil {
// 		orderByIn(root.Left, root)
// 	}

// }

// func subtreeSum(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return root.Val + subtreeSum(root.Right) + subtreeSum(root.Left)
// }

// Runtime: 196 ms, faster than 97.16% of Go online submissions for Convert BST to Greater Tree.
// Memory Usage: 272.6 MB, less than 40.65% of Go online submissions for Convert BST to Greater Tree.

// 用反转中序遍历
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := make([]*TreeNode, 0)
	orderByIn(root, &res)

	for i := 1; i < len(res); i++ {
		res[i].Val += res[i-1].Val
	}

	return root
}

func orderByIn(root *TreeNode, res *[]*TreeNode) {
	if root.Right != nil {
		orderByIn(root.Right, res)
	}

	*res = append(*res, root)

	if root.Left != nil {
		orderByIn(root.Left, res)
	}
}
