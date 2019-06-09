package uniquePaths

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?


Above is a 7 x 3 grid. How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right

Example 2:

Input: m = 7, n = 3
Output: 28


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
// Memory Usage: 1.9 MB, less than 93.81% of Go online submissions for Unique Paths.

func uniquePaths(m int, n int) int {
	total := m + n - 2
	if n-1 > total/2 {
		n = total - (n - 1) + 1
	}
	return combination(total, n-1)
}

func combination(n, m int) int {
	bigger := 1
	smaller := 1
	for i := 1; i <= m; i++ {
		bigger *= n
		smaller *= i
		n--
	}
	return bigger / smaller
}
