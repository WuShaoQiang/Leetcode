package containsNearbyDuplicate

/*
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

*/

// Runtime: 16 ms, faster than 82.35% of Go online submissions for Contains Duplicate II.
// Memory Usage: 9 MB, less than 23.92% of Go online submissions for Contains Duplicate II.

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for idx, num := range nums {
		if dis, exist := m[num]; exist && idx-dis <= k {
			return true
		}
		m[num] = idx
	}
	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums) && j <= i+k; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
