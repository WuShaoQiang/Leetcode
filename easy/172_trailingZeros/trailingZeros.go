package trailingZeros

/*
Given an integer n, return the number of trailing zeroes in n!.

Example 1:

Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:

Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.

Note: Your solution should be in logarithmic time complexity.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2 MB, less than 90.98% of Go online submissions for Factorial Trailing Zeroes.

func trailingZeroes(n int) int {
	count := 0
	tmp := 1
	len := 0
	for n/(tmp*5) > 0 {
		len++
		tmp = tmp * 5
	}
	for i := 0; i < len; i++ {
		count += (n / tmp)
		tmp = tmp / 5
	}
	return count
}
