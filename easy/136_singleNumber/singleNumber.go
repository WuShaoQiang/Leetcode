package singleNumber

/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1

Example 2:

Input: [4,1,2,1,2]
Output: 4

*/

// Runtime: 8 ms, faster than 99.41% of Go online submissions for Single Number.
// Memory Usage: 4.7 MB, less than 71.04% of Go online submissions for Single Number.

func singleNumber(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		nums[i-1] = nums[i-1] ^ nums[i]
	}
	return nums[0]
}
