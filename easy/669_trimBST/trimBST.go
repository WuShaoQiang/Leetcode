package trimBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:

Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2

Example 2:

Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
*/

// 执行用时 : 32 ms, 在Trim a Binary Search Tree的Go提交中击败了54.84% 的用户
// 内存消耗 : 6.7 MB, 在Trim a Binary Search Tree的Go提交中击败了66.67% 的用户

// func trimBST(root *TreeNode, L int, R int) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	// 选出根节点
// 	for {
// 		// 根节点大于最大值,所以要取其某一个左子树
// 		for root.Val > R && root != nil {
// 			root = root.Left
// 		}

// 		for root.Val < L && root != nil {
// 			root = root.Right
// 		}

// 		if root.Val >= L && root.Val <= R {
// 			break
// 		}
// 	}

// 	//选出根节点了

// 	// 判断临界
// 	// root是最小值,只有右边了
// 	if root.Val == L {
// 		root.Left = nil
// 	}
// 	// root是最大值,只有左边了
// 	if root.Val == R {
// 		root.Right = nil
// 	}

// 	tmp := root
// 	// left
// 	if tmp.Left != nil {
// 		for tmp != nil {
// 			// 叶子结点,已经没有要选择指向的结点了
// 			if tmp.Left == nil && tmp.Right == nil {
// 				break
// 			}

// 			// 如果左孩子小于L,那么抛弃这个结点,连接这个左孩子的右孩子
// 			if tmp.Left != nil && tmp.Left.Val < L {
// 				tmp.Left = tmp.Left.Right
// 			} else if tmp.Left != nil {
// 				tmp = tmp.Left
// 				// 如果左孩子为nil,那么就选右孩子进行传递
// 			} else {
// 				tmp = tmp.Right
// 			}
// 		}
// 	}

// 	tmp = root
// 	// right
// 	if tmp.Right != nil {
// 		for tmp != nil {
// 			// 叶子结点,已经没有要选择指向的结点了
// 			if tmp.Left == nil && tmp.Right == nil {
// 				break
// 			}

// 			if tmp.Right != nil && tmp.Right.Val > R {
// 				tmp.Right = tmp.Right.Left
// 			} else if tmp.Right != nil {
// 				tmp = tmp.Right
// 			} else {
// 				tmp = tmp.Left
// 			}
// 		}
// 	}

// 	return root
// }

// 递归做法

// 执行用时 : 24 ms, 在Trim a Binary Search Tree的Go提交中击败了96.77% 的用户
// 内存消耗 : 6.9 MB, 在Trim a Binary Search Tree的Go提交中击败了53.33% 的用户

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	// 相当于忽视了这个root.Val
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	// 相当于忽视了这个root.Val
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	// 不可忽视的Val
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
