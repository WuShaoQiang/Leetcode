package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := (len(nums) - 1) / 2

	head := &TreeNode{Val: nums[mid]}
	heightHalancedTree(head, nums, 0, mid-1)
	heightHalancedTree(head, nums, mid+1, len(nums)-1)

	return head
}

func heightHalancedTree(t *TreeNode, nums []int, left, right int) {
	if left <= right {
		mid := (left + right) / 2
		// fmt.Println("insert:", nums[mid])
		insert(t, nums[mid])
		heightHalancedTree(t, nums, left, mid-1)
		heightHalancedTree(t, nums, mid+1, right)
	}
}

func insert(t *TreeNode, num int) {
	tmp := t
	for {
		//right
		if num > tmp.Val {
			// create
			if tmp.Right == nil {
				tmp.Right = &TreeNode{Val: num}
				return
			}

			// find next
			tmp = tmp.Right
		}

		// left
		if num <= tmp.Val {
			// create
			if tmp.Left == nil {
				tmp.Left = &TreeNode{Val: num}
				return
			}
			// find next
			tmp = tmp.Left
		}
	}
}
