package increasingBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9

Note:

    The number of nodes in the given tree will be between 1 and 100.
    Each node will have a unique integer value from 0 to 1000.

*/

// Runtime: 24 ms, faster than 96.47% of Go online submissions for Increasing Order Search Tree.
// Memory Usage: 6.4 MB, less than 32.35% of Go online submissions for Increasing Order Search Tree.

// 中序遍历,这种是在原来的点上做变化
// func increasingBST(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	res := make([]*TreeNode, 0)
// 	inOrder(root, &res)
// 	for i := 0; i < len(res)-1; i++ {
// 		res[i].Right = res[i+1]
// 		res[i].Left = nil
// 	}

// 	res[len(res)-1].Right = nil
// 	res[len(res)-1].Left = nil

// 	return res[0]
// }

// func inOrder(root *TreeNode, res *[]*TreeNode) {
// 	if root.Left != nil {
// 		inOrder(root.Left, res)
// 	}

// 	*res = append(*res, root)

// 	if root.Right != nil {
// 		inOrder(root.Right, res)
// 	}

// }

// Runtime: 12 ms, faster than 100.00% of Go online submissions for Increasing Order Search Tree.
// Memory Usage: 6.4 MB, less than 48.53% of Go online submissions for Increasing Order Search Tree.

// 这种就是新创建点
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	inOrder(root, &res)
	head := &TreeNode{Val: res[0]}
	tmp := head
	for i := 1; i < len(res); i++ {
		tmp.Right = &TreeNode{Val: res[i]}
		tmp = tmp.Right
	}

	return head
}

func inOrder(root *TreeNode, res *[]int) {
	if root.Left != nil {
		inOrder(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		inOrder(root.Right, res)
	}
}
