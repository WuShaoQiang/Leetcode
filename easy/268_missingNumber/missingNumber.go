package missingNumber

/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2

Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?*/

// Runtime: 16 ms, faster than 99.04% of Go online submissions for Missing Number.
// Memory Usage: 5.9 MB, less than 86.87% of Go online submissions for Missing Number.

// 根据没丢失的和减去丢失数组的和，得到的就是丢失的元素
func missingNumber(nums []int) int {
	total := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		total += i
		sum += nums[i]
	}
	total += len(nums)

	return total - sum
}
