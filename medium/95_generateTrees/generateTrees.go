package generateTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

// Runtime: 20 ms, faster than 96.48% of Go online submissions for Unique Binary Search Trees II.
// Memory Usage: 41 MB, less than 42.62% of Go online submissions for Unique Binary Search Trees II.

// 这是看参考写出来的
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return genTreeList(1, n)
}

func genTreeList(start, end int) []*TreeNode {
	list := make([]*TreeNode, 0)

	if start > end {
		list = append(list, nil)
	}

	for i := start; i <= end; i++ {
		leftList := genTreeList(start, i-1)
		rightList := genTreeList(i+1, end)
		for _, left := range leftList {
			for _, right := range rightList {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				list = append(list, root)
			}
		}
	}

	return list
}
