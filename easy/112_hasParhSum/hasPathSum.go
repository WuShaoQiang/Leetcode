package hasParhSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return count(root, sum)
}

func count(root *TreeNode, sum int) bool {
	if root != nil {
		sum = sum - root.Val
		// 右边为空，左边有值并不是叶子节点
		if root.Right == nil && root.Left != nil {
			return count(root.Left, sum)
		}
		// 左边为空，右边有值也不是叶子节点
		if root.Right != nil && root.Left == nil {
			return count(root.Right, sum)
		}
		return count(root.Left, sum) || count(root.Right, sum)
	}
	//必须是叶子节点
	if root == nil {
		if sum == 0 {
			return true
		}
	}
	return false
}
