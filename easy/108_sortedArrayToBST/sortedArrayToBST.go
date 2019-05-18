package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

*/

// Runtime: 100 ms, faster than 97.58% of Go online submissions for Convert Sorted Array to Binary Search Tree.
// Memory Usage: 273.4 MB, less than 5.11% of Go online submissions for Convert Sorted Array to Binary Search Tree.

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
