package longestUnivaluePath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

The length of path between two nodes is represented by the number of edges between them.



Example 1:

Input:

              5
             / \
            4   5
           / \   \
          1   1   5

Output: 2



Example 2:

Input:

              1
             / \
            4   5
           / \   \
          4   4   5

Output: 2



Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

*/

// Runtime: 68 ms, faster than 100.00% of Go online submissions for Longest Univalue Path.
// Memory Usage: 6.9 MB, less than 82.14% of Go online submissions for Longest Univalue Path.

// 与第543题类似,这道题是参考别人的答案做出来的
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxH := 0
	getHeight(root, &maxH)
	return maxH
}

func getHeight(root *TreeNode, maxH *int) int {
	if root == nil {
		return 0
	}
	tmp := 0
	left := getHeight(root.Left, maxH)
	right := getHeight(root.Right, maxH)

	// 如果这个节点两个孩子的值都与它相等,那么有机会以这个节点为根节点的路径是最长的
	if root.Left != nil && root.Right != nil && root.Left.Val == root.Val && root.Right.Val == root.Val {
		*maxH = max(*maxH, left+right+2)
	}

	// 就算某个节点有双子树,但是也不能因此对上层停止传递,所以还是要将这个节点单独的一边传上去的
	if root.Left != nil && root.Left.Val == root.Val {
		tmp = left + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		tmp = max(tmp, right+1)
	}

	*maxH = max(*maxH, tmp)
	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
