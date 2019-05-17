package climbStairs

import "math"

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
// Memory Usage: 1.9 MB, less than 88.83% of Go online submissions for Climbing Stairs.

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	count2 := n / 2
	result := 0
	for ; count2 >= 0; count2-- {
		result += calc(n-count2, count2)
	}
	return result
}

//注意float的运算可能会出现一些.9999或者.00001的情况
func calc(m, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return m
	}
	if n < (m+1)/2 {
		n = m - n
	}
	tmp := 1.0
	for i := m; i > n; i-- {
		tmp = ((tmp * float64(i)) / (float64(i) - float64(n)))
	}
	return int(math.Round(tmp))
}
