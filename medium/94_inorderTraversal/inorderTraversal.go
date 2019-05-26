package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]

Follow up: Recursive solution is trivial, could you do it iteratively?

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2.1 MB, less than 24.30% of Go online submissions for Binary Tree Inorder Traversal.

// 递归方法很简单
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	res := make([]int, 0)
// 	if root.Left != nil {
// 		res = append(res, inorderTraversal(root.Left)...)
// 	}

// 	res = append(res, root.Val)

// 	if root.Right != nil {
// 		res = append(res, inorderTraversal(root.Right)...)
// 	}

// 	return res
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 2.1 MB, less than 72.71% of Go online submissions for Binary Tree Inorder Traversal.

// 迭代 (参考)

func inorderTraversal(root *TreeNode) []int {
	r := []int{}

	// 模拟栈存储节点
	s := []*TreeNode{}
	c := root
	for c != nil || len(s) > 0 {
		if c != nil {
			// 如果不为nil,则存入栈中
			s = append(s, c)
			c = c.Left
		} else {
			// 如果为nil,没有左子树了,取出第一个节点,并且删除
			p := s[len(s)-1]
			s = s[:len(s)-1]
			// 将该节点是值加入到数组
			r = append(r, p.Val)
			// 再去寻找右子树
			c = p.Right
		}
	}
	return r
}
