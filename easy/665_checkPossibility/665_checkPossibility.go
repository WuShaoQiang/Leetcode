package checkPossibility

/*
 Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:

Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

Example 2:

Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.

Note: The n belongs to [1, 10,000].
*/

// Runtime: 24 ms, faster than 93.86% of Go online submissions for Non-decreasing Array.
// Memory Usage: 6.2 MB, less than 10.77% of Go online submissions for Non-decreasing Array.

func checkPossibility(nums []int) bool {
	var i int
	for ; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i != 0 {
				tmp1 := append([]int{}, nums...)
				tmp2 := append([]int{}, nums...)
				tmp1[i] = tmp1[i-1]
				tmp2[i+1] = tmp2[i]
				return check(tmp1) || check(tmp2)
			} else {
				return check(nums[1:])
			}
		}
	}
	return true

}

func check(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

// Runtime: 24 ms, faster than 93.86% of Go online submissions for Non-decreasing Array.
// Memory Usage: 6.2 MB, less than 96.92% of Go online submissions for Non-decreasing Array.

// func checkPossibility(nums []int) bool {
// 	n := len(nums)
// 	p := -1
// 	for i := 0; i < n-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			if p >= 0 {
// 				return false
// 			}
// 			p = i
// 		}
// 	}
// 	// 如果是在头或者尾，那可以任意改动，所以直接返回true
// 	if p <= 0 || p == n-2 {
// 		return true
// 	}

// 	// 到这里的前提是，有一个数p > 后面的数p+1，并且整个数组只发生过一次递减情况
// 	// 如果这个数 前面的数<=后面的数 说明后面递增的数不会比前面的小,递减关系只有一处
// 	// 如果这个数 <= 后面第二个数，说明递减关系只有一处
// 	if nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2] {
// 		return true
// 	}
// 	return false
// }
