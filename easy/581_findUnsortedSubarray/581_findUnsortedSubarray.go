package findUnsortedSubarray

import "math"

/*
Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:

Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

Note:

    Then length of the input array is in range [1, 10,000].
    The input array may contain duplicates, so ascending order here means <=.

*/

// func findUnsortedSubarray(nums []int) int {
// 	res := len(nums)
// 	var i int
// 	for ; i < len(nums); i++ {
// 		if !isMin(nums[i], nums[i:]) {
// 			break
// 		}
// 		res--
// 	}

// 	for j := len(nums) - 1; j >= i; j-- {
// 		if !isMax(nums[j], nums[i:j+1]) {
// 			break
// 		}
// 		res--
// 	}

// 	return res

// }

// func isMin(num int, nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		if num > nums[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isMax(num int, nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		if num < nums[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// Runtime: 24 ms, faster than 99.25% of Go online submissions for Shortest Unsorted Continuous Subarray.
// Memory Usage: 6.2 MB, less than 60.55% of Go online submissions for Shortest Unsorted Continuous Subarray.

func findUnsortedSubarray(nums []int) int {
	start, end := -1, -1
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		// 每次找到前面的比后面的大，都要再次确认start的位置
		// 这里进去一般i都为1了
		if max > nums[i] {
			if start == -1 {
				start = i - 1
			}
			// 这里保证开头的值是最小的或者是最前的
			for start > 0 && nums[start-1] > nums[i] {
				start--
			}
			// 这里加一是因为索引双边都算，单单一个减法只算一边，比如：2-1=1，但是我包含的应该是2和1，所以应该是得到2
			end = i + 1
		} else {
			max = nums[i]
		}
	}
	return end - start
}
