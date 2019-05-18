package thirdMax

import "math"

/*
Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:

Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.

Example 2:

Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:

Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Third Maximum Number.
// Memory Usage: 3.9 MB, less than 10.74% of Go online submissions for Third Maximum Number.

func thirdMax(nums []int) int {
	max := make([]int, 3)
	for i := 0; i < 3; i++ {
		max[i] = math.MinInt64
	}
	count := 0
	isExist := make(map[int]bool)
	for _, num := range nums {
		if _, exist := isExist[num]; !exist {
			count++
			isExist[num] = true
			if num > max[0] {
				copy(max[1:3], max[0:2])
				max[0] = num
			} else if num > max[1] {
				max[2] = max[1]
				max[1] = num
			} else if num > max[2] {
				max[2] = num
			}
		}
	}

	if count < 3 {
		return max[0]
	}

	return max[2]
}

// 执行用时 : 8 ms, 在Third Maximum Number的Go提交中击败了91.67% 的用户
// 内存消耗 : 4 MB, 在Third Maximum Number的Go提交中击败了23.33% 的用户

// func thirdMax(nums []int) int {
// 	max := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}
// 	count := 0
// 	isExist := make(map[int]bool)
// 	for _, num := range nums {
// 		if _, exist := isExist[num]; !exist {
// 			count++
// 			isExist[num] = true
// 			if num > max[0] {
// 				copy(max[1:3], max[0:2])
// 				max[0] = num
// 			} else if num > max[1] {
// 				max[2] = max[1]
// 				max[1] = num
// 			} else if num > max[2] {
// 				max[2] = num
// 			}
// 		}
// 	}

// 	if count < 3 {
// 		return max[0]
// 	}

// 	return max[2]
// }
