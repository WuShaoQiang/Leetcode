package twoSum

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

    Your returned answers (both index1 and index2) are not zero-based.
    You may assume that each input would have exactly one solution and you may not use the same element twice.

Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

*/

// Runtime: 4 ms, faster than 99.24% of Go online submissions for Two Sum II - Input array is sorted.
// Memory Usage: 3 MB, less than 56.16% of Go online submissions for Two Sum II - Input array is sorted.

// 左右夹
func twoSum(numbers []int, target int) []int {
	idx1 := 0
	idx2 := len(numbers) - 1

	// for target < numbers[idx2] {
	// 	idx2--
	// }
	for {
		if idx1 == idx2 {
			return nil
		}
		if result := numbers[idx1] + numbers[idx2]; result != target {
			if result < target {
				idx1++
			} else {
				idx2--
			}
		} else {
			break
		}
	}
	return []int{idx1 + 1, idx2 + 1}
}
