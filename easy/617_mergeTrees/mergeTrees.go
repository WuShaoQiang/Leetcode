package mergeTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

注意: 合并必须从两个树的根节点开始。

*/

// 执行用时 : 44 ms, 在Merge Two Binary Trees的Go提交中击败了96.46% 的用户
// 内存消耗 : 8.6 MB, 在Merge Two Binary Trees的Go提交中击败了46.94% 的用户

// 第一个想法就是同时遍历两颗树
// func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
// 	if t1 == nil {
// 		return t2
// 	}

// 	if t2 == nil {
// 		return t1
// 	}

// 	// 都不为nil
// 	iteration(t1, t2)

// 	return t1
// }

// func iteration(t1 *TreeNode, t2 *TreeNode) {
// 	// 如果两棵树对应节点都不为nil,直接在t1上的值改变
// 	t1.Val += t2.Val

// 	// 双方左子树都存在
// 	if t1.Left != nil && t2.Left != nil {
// 		iteration(t1.Left, t2.Left)
// 	} else if t1.Left == nil && t2.Left != nil { // 需要直接指向的情况
// 		t1.Left = t2.Left
// 	}
// 	// 双方右子树存在
// 	if t1.Right != nil && t2.Right != nil {
// 		iteration(t1.Right, t2.Right)
// 	} else if t1.Right == nil && t2.Right != nil { // 需要直接指向的情况
// 		t1.Right = t2.Right
// 	}

// 	// 剩下的情况不用处理,比如双方nil,或者t1!=nil&&t2==nil
// 	return
// }

// 递归的做法
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
		return t1
	} else if t1 == nil && t2 != nil {
		return t2
	} else {
		return t1
	}
}
