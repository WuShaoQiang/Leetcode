package pivotIndex

/*
Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define the pivot index as the index where the sum of the numbers to the left of the index is equal to the sum of the numbers to the right of the index.

If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.

Example 1:

Input:
nums = [1, 7, 3, 6, 5, 6]
Output: 3
Explanation:
The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.



Example 2:

Input:
nums = [1, 2, 3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.



Note:

    The length of nums will be in the range [0, 10000].
    Each element nums[i] will be an integer in the range [-1000, 1000].

*/

// Runtime: 20 ms, faster than 95.11% of Go online submissions for Find Pivot Index.
// Memory Usage: 5.9 MB, less than 67.62% of Go online submissions for Find Pivot Index.

// 想法是将某一个节点的左右两边加起来做比较，可以转换为：左边从0开始，右边从最大的和开始，这样子方便算某一个节点的左右两边
func pivotIndex(nums []int) int {
	rightSum, leftSum := 0, 0
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}

	// 这个i就相当于一个中间界
	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1

}
