package containDuplicate

/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true

Example 2:

Input: [1,2,3,4]
Output: false

Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true

*/

// Runtime: 20 ms, faster than 95.18% of Go online submissions for Contains Duplicate.
// Memory Usage: 9.2 MB, less than 49.13% of Go online submissions for Contains Duplicate.

func containsDuplicate(nums []int) bool {
	isExist := make(map[int]bool)

	for _, num := range nums {
		if _, exist := isExist[num]; exist {
			return true
		}

		isExist[num] = true
	}
	return false
}
