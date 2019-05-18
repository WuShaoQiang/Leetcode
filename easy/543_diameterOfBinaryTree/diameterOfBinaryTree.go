package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

// 示例 :
// 给定二叉树

//           1
//          / \
//         2   3
//        / \
//       4   5

// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

// 注意：两结点之间的路径长度是以它们之间边的数目表示。

// 执行用时 : 28 ms, 在Diameter of Binary Tree的Go提交中击败了11.54% 的用户
// 内存消耗 : 4.5 MB, 在Diameter of Binary Tree的Go提交中击败了95.83% 的用户

// func diameterOfBinaryTree(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	var max int

// 	iteration(root, &max)
// 	return max
// }

// func iteration(root *TreeNode, max *int) {
// 	if root != nil {
// 		tmp := diameterOfNode(root)
// 		if *max < tmp {
// 			*max = tmp
// 		}
// 		iteration(root.Left, max)
// 		iteration(root.Right, max)
// 	}
// }

// func diameterOfNode(root *TreeNode) int {
// 	var left, right int
// 	if root.Left == nil {
// 		left = 0
// 	} else {
// 		left = 1 + depthOfNode(root.Left)
// 	}

// 	if root.Right == nil {
// 		right = 0
// 	} else {
// 		right = 1 + depthOfNode(root.Right)
// 	}

// 	return left + right
// }

// func depthOfNode(root *TreeNode) int {

// 	if root.Right != nil && root.Left != nil {
// 		return max(1+depthOfNode(root.Left), 1+depthOfNode(root.Right))
// 	}

// 	if root.Right != nil {
// 		return 1 + depthOfNode(root.Right)
// 	}

// 	if root.Left != nil {
// 		return 1 + depthOfNode(root.Left)
// 	}

// 	return 0
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }

// Runtime: 4 ms, faster than 99.47% of Go online submissions for Diameter of Binary Tree.
// Memory Usage: 4.5 MB, less than 77.18% of Go online submissions for Diameter of Binary Tree.

func diameterOfBinaryTree(root *TreeNode) int {
	var maxH int
	getHeight(root, &maxH)

	return maxH
}

func getHeight(root *TreeNode, maxH *int) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left, maxH)
	right := getHeight(root.Right, maxH)
	*maxH = max(*maxH, left+right)
	return max(right, left) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
