package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}

// 	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
// 		return false
// 	}

// 	return isSym(append(make([]*TreeNode, 0), root.Left, root.Right), false)
// }

// func isSym(nodes []*TreeNode, null bool) bool {
// 	if !null {
// 		allNull := true
// 		n := len(nodes)
// 		if n%2 != 0 {
// 			return false
// 		}
// 		nextNodes := make([]*TreeNode, 2*n)

// 		for i := 0; i < n/2; i++ {
// 			if nodes[i] == nil || nodes[n-i-1] == nil {
// 				if !(nodes[i] == nil && nodes[n-i-1] == nil) {
// 					return false
// 				}
// 			} else if nodes[i].Val != nodes[n-i-1].Val {
// 				return false
// 			} else {
// 				// fmt.Printf("node %v value %v = node %v value %v\n", i, nodes[i].Val, n-i-1, nodes[n-i-1].Val)
// 				nextNodes[2*i], nextNodes[2*i+1] = nodes[i].Left, nodes[i].Right
// 				nextNodes[2*(n-i-1)], nextNodes[2*(n-i-1)+1] = nodes[n-i-1].Left, nodes[n-i-1].Right
// 				if nodes[i].Left != nil || nodes[i].Right != nil || nodes[n-i-1].Left != nil || nodes[n-i-1].Right != nil {
// 					allNull = false
// 				}
// 			}
// 			// fmt.Printf("i : %v, n : %v\n", i, n)

// 		}
// 		return isSym(nextNodes, allNull)
// 	}
// 	return true
// }
