package isCousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.



Example 1:

Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Example 2:

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true

Example 3:

Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false



Note:

    The number of nodes in the tree will be between 2 and 100.
    Each node has a unique integer value from 1 to 100.

*/

// Runtime: 12 ms, faster than 9.40% of Go online submissions for Cousins in Binary Tree.
// Memory Usage: 8.1 MB, less than 5.34% of Go online submissions for Cousins in Binary Tree.

//要有同样的深度,但是一定要不同parent
// func isCousins(root *TreeNode, x int, y int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	queue := []*TreeNode{root}

// 	for {
// 		l := len(queue)
// 		if l == 0 {
// 			break
// 		}
// 		// 如果count=2,说明两个数在同一层
// 		count := 0
// 		for l > 0 {
// 			tmp := queue[l-1]
// 			isSameParent := 0
// 			if tmp.Left != nil {
// 				queue = append(queue, tmp.Left)
// 				if tmp.Left.Val == x || tmp.Left.Val == y {
// 					isSameParent++
// 					count++
// 				}
// 			}

// 			if tmp.Right != nil {
// 				queue = append(queue, tmp.Right)
// 				if tmp.Right.Val == x || tmp.Right.Val == y {
// 					isSameParent++
// 					count++
// 				}
// 			}
// 			// 如果等于二,说明是同一个parent
// 			if isSameParent == 2 {
// 				return false
// 			}
// 			l--
// 		}
// 		// 说明这一层只有一个
// 		if count == 1 {
// 			return false
// 		} else if count == 2 {
// 			return true
// 		}
// 	}
// 	return false
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Cousins in Binary Tree.
// Memory Usage: 2.7 MB, less than 7.63% of Go online submissions for Cousins in Binary Tree.

// 用map记录
func isCousins(root *TreeNode, x int, y int) bool {
	depthMap, parentMap := make(map[int]int), make(map[int]int)
	dfs(root, 101, depthMap, parentMap)
	return depthMap[x] == depthMap[y] && parentMap[x] != parentMap[y]
}

func dfs(node *TreeNode, parent int, depthMap map[int]int, parentMap map[int]int) {
	if node != nil {
		if parent == 101 {
			depthMap[node.Val] = 0
		} else {
			depthMap[node.Val] = 1 + depthMap[parent]
		}
		parentMap[node.Val] = parent
		dfs(node.Left, node.Val, depthMap, parentMap)
		dfs(node.Right, node.Val, depthMap, parentMap)
	}
}
