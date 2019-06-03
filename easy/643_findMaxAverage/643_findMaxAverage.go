package findMaxAverage

import (
	"math"
)

/*
Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value. And you need to output the maximum average value.

Example 1:

Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75



Note:

    1 <= k <= n <= 30,000.
    Elements of the given array will be in the range [-10,000, 10,000].

*/

// Runtime: 124 ms, faster than 98.72% of Go online submissions for Maximum Average Subarray I.
// Memory Usage: 7.6 MB, less than 47.92% of Go online submissions for Maximum Average Subarray I.

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	left := 0
	right := left + k - 1
	max := math.MinInt64
	sum := 0
	for i := left; i <= right; i++ {
		sum += nums[i]
	}

	if sum > max {
		max = sum
	}
	right++

	for right < len(nums) {
		sum += nums[right]
		sum -= nums[left]
		if sum > max {
			max = sum
		}
		right++
		left++
	}

	return float64(max) / float64(k)
}

// Runtime: 116 ms, faster than 100.00% of Go online submissions for Maximum Average Subarray I.
// Memory Usage: 7.7 MB, less than 22.92% of Go online submissions for Maximum Average Subarray I.

// func findMaxAverage(nums []int, k int) float64 {
// 	s := 0
// 	for i := 0; i < k; i++ {
// 		s = s + nums[i]
// 	}
// 	max_sum := s
// 	for i := k; i < len(nums); i++ {
// 		s = s - nums[i-k] + nums[i]
// 		if s > max_sum {
// 			max_sum = s
// 		}
// 	}
// 	return float64(max_sum) / float64(k)
// }
