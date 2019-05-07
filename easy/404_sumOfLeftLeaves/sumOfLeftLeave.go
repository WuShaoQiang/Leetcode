package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	search(root, false, &sum)
	return sum

}

func search(root *TreeNode, flag bool, sum *int) {
	if root == nil {
		return
	}

	if flag && root.Left == nil && root.Right == nil {
		*sum += root.Val
	}

	if root.Left != nil {
		search(root.Left, true, sum)
	}

	if root.Right != nil {
		search(root.Right, false, sum)
	}
}
