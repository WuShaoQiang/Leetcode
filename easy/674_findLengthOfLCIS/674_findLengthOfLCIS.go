package findLengthOfLCIS

/*
 Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:

Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.

Example 2:

Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.

Note: Length of the array will not exceed 10,000.
*/

// Runtime: 8 ms, faster than 98.35% of Go online submissions for Longest Continuous Increasing Subsequence.
// Memory Usage: 4.5 MB, less than 62.67% of Go online submissions for Longest Continuous Increasing Subsequence.

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	count := 1
	max := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 1
		}
	}
	if max < count {
		max = count
	}
	return max
}
