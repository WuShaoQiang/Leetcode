package minPathSum

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.


*/

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Minimum Path Sum.
// Memory Usage: 3.9 MB, less than 78.41% of Go online submissions for Minimum Path Sum.

// 和63题的做法一样，都是累计的动态规划题
func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
