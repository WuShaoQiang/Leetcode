package findMaxConsecutiveOnes

/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:

Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.

Note:

    The input array will only contain 0 and 1.
    The length of input array is a positive integer and will not exceed 10,000


*/

// Runtime: 36 ms, faster than 96.99% of Go online submissions for Max Consecutive Ones.
// Memory Usage: 6.3 MB, less than 55.07% of Go online submissions for Max Consecutive Ones.

// func findMaxConsecutiveOnes(nums []int) int {
// 	max := 0
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			if count > max {
// 				max = count
// 			}
// 			count = 0
// 		} else {
// 			count++
// 		}
// 	}
// 	if count > max {
// 		max = count
// 	}
// 	return max
// }
