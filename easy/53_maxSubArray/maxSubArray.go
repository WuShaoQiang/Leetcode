package maxSubArray

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Maximum Subarray.
// Memory Usage: 3.3 MB, less than 89.89% of Go online submissions for Maximum Subarray.

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	total := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果发现加起来是个负数,那么舍弃前面的
		if total < 0 {
			total = 0
		}
		total += nums[i]
		// 每次加一个都要计算最大值
		if max < total {
			max = total
		}
	}
	return max
}
