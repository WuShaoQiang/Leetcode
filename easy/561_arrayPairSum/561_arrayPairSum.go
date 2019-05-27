package arrayPairSum

import (
	"sort"
)

/*
 Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:

Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).

Note:

    n is a positive integer, which is in the range of [1, 10000].
    All the integers in the array will be in the range of [-10000, 10000].

*/

// Runtime: 72 ms, faster than 77.63% of Go online submissions for Array Partition I.
// Memory Usage: 6.7 MB, less than 92.98% of Go online submissions for Array Partition I.

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := len(nums) - 2; i >= 0; i -= 2 {
		sum += nums[i]
	}
	return sum
}

// Runtime: 40 ms, faster than 100.00% of Go online submissions for Array Partition I.
// Memory Usage: 6.8 MB, less than 19.30% of Go online submissions for Array Partition I.

// func arrayPairSum(nums []int) int {
//     var buckets [20001]int8
// 	for _, num := range nums {
// 		buckets[num+10000]++
// 	}

// 	sum := 0
// 	needAdd := true
// 	for num, count := range buckets {
// 		for count > 0 {
// 			if needAdd {
// 				sum += num - 10000
// 			}
// 			needAdd = !needAdd
// 			count--
// 		}
// 	}

// 	return sum
// }
