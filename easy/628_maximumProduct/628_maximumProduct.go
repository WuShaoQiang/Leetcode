package maximumProduct

import (
	"math"
)

/*
Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:

Input: [1,2,3]
Output: 6



Example 2:

Input: [1,2,3,4]
Output: 24



Note:

    The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
    Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.

*/

// Runtime: 36 ms, faster than 99.00% of Go online submissions for Maximum Product of Three Numbers.
// Memory Usage: 6.3 MB, less than 76.00% of Go online submissions for Maximum Product of Three Numbers.

func maximumProduct(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64+2, math.MinInt64+1, math.MinInt64
	firstMin, secondMin := math.MaxInt64-1, math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if firstMax < nums[i] {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = nums[i]
		} else if secondMax < nums[i] {
			thirdMax = secondMax
			secondMax = nums[i]
		} else if thirdMax < nums[i] {
			thirdMax = nums[i]
		}

		if firstMin > nums[i] {
			secondMin = firstMin
			firstMin = nums[i]
		} else if secondMin > nums[i] {
			secondMin = nums[i]
		}
	}

	res1 := firstMax * firstMin * secondMin
	res2 := firstMax * secondMax * thirdMax

	if res1 > res2 {
		return res1
	}

	return res2

}
