package judgeSquareSum

import "math"

/*
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:

Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5



Example 2:

Input: 3
Output: False

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Square Numbers.
// Memory Usage: 2 MB, less than 12.50% of Go online submissions for Sum of Square Numbers.

func judgeSquareSum(c int) bool {
	max := int(math.Sqrt(float64(c)))
	min := 0
	for min <= max {
		res := max*max + min*min
		if res == c {
			return true
		} else if res > c {
			max--
		} else {
			min++
		}
	}
	return false
}
