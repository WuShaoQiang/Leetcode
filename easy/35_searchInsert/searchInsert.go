package searchInsert

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2

Example 2:

Input: [1,3,5,6], 2
Output: 1

Example 3:

Input: [1,3,5,6], 7
Output: 4

Example 4:

Input: [1,3,5,6], 0
Output: 0

*/

// Runtime: 4 ms, faster than 98.92% of Go online submissions for Search Insert Position.
// Memory Usage: 3.1 MB, less than 49.90% of Go online submissions for Search Insert Position.

// func searchInsert(nums []int, target int) int {
// 	isFound, idx := binarySearch(nums, 0, len(nums)-1, target)
// 	if isFound {
// 		return idx
// 	}

// 	if nums[idx] > target {
// 		return idx
// 	} else {
// 		return idx + 1
// 	}
// }

// func binarySearch(nums []int, left, right, target int) (found bool, idx int) {
// 	if left < right {
// 		mid := (left + right) / 2

// 		if nums[mid] == target {
// 			return true, mid
// 		} else if nums[mid] > target {
// 			return binarySearch(nums, left, mid-1, target)
// 		} else {
// 			return binarySearch(nums, mid+1, right, target)
// 		}
// 	} else if nums[left] == target {
// 		return true, left
// 	} else {
// 		return false, left
// 	}
// }

// Runtime: 4 ms, faster than 98.92% of Go online submissions for Search Insert Position.
// Memory Usage: 3.1 MB, less than 49.90% of Go online submissions for Search Insert Position.

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 没找到
	return left

}
