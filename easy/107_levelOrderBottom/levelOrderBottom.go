package levelOrderBottom

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return nil
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		len := len(que)
		var level []int
		for i := 0; i < len; i++ {
			node := que[0]
			que = que[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		ret = append(ret, level)
	}
	var ans [][]int
	for i := len(ret) - 1; i >= 0; i-- {
		ans = append(ans, ret[i])
	}
	return ans
}
