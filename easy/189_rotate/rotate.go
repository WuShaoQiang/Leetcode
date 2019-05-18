package rotate

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Note:

    Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
    Could you do it in-place with O(1) extra space?
*/

// Runtime: 44 ms, faster than 100.00% of Go online submissions for Rotate Array.
// Memory Usage: 7.6 MB, less than 48.53% of Go online submissions for Rotate Array.

func rotate(nums []int, k int) {
	// 小于等于0直接返回
	if k <= 0 {
		return
	}
	// 如果大于长度，取余数部分再传递一次，记得return
	if k > len(nums) {
		rotate(nums, k%len(nums))
		return
	}

	// 和长度一样等于没变
	if k == len(nums) {
		return
	}
	// 如果数组长度不足以支持翻转，直接返回
	if len(nums) == 1 || len(nums) == 0 {
		return
	}

	// 计算出从哪里分隔
	idx := len(nums) - k
	reverse(nums[:idx])
	reverse(nums[idx:])
	reverse(nums)
}

func reverse(nums []int) {
	l := len(nums)
	// 如果长度为1，不用翻转
	if l == 1 {
		return
	}
	n := (l - 1) / 2
	for n >= 0 {
		nums[n], nums[l-1-n] = nums[l-1-n], nums[n]
		n--
	}
}

// func rotate(nums []int, k int) {
// 	l := len(nums)
// 	for i := l - 1; i > l-1-k; i-- {
// 		tmp := nums[l-1]
// 		for j := l - 1; j > 0; j-- {
// 			nums[j] = nums[j-1]
// 		}
// 		nums[0] = tmp
// 	}
// }
