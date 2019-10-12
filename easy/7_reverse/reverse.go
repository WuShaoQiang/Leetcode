package reverse

import "math"

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 87.10% of Go online submissions for Reverse Integer.

// -2147483648~2147483647
func reverse(x int) int {
	var rev int
	for x != 0 {
		lastNum := x % 10
		x = x / 10
		if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && lastNum > 7 {
			return 0
		}
		if rev < math.MinInt32/10 || rev == math.MinInt32/10 && lastNum < -8 {
			return 0
		}
		rev = rev*10 + lastNum
	}
	return rev
}
