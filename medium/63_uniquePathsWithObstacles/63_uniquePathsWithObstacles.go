package uniquePathsWithObstacles

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right


*/

// Time Limit Exceeded

// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	res := 0
// 	if obstacleGrid[0][0] == 1 {
// 		return 0
// 	}
// 	Go(obstacleGrid, 0, 0, &res)
// 	return res
// }

// func Go(grid [][]int, currRow, currCol int, res *int) {
// 	//到达终点
// 	if currRow == len(grid)-1 && currCol == len(grid[0])-1 {
// 		*res++
// 		return
// 	}
// 	// 如果右边还有位置，并且不为1
// 	if currCol < len(grid[0])-1 && grid[currRow][currCol+1] != 1 {
// 		Go(grid, currRow, currCol+1, res)
// 	}
// 	// 如果下面还有位置，并且不为1
// 	if currRow < len(grid)-1 && grid[currRow+1][currCol] != 1 {
// 		Go(grid, currRow+1, currCol, res)
// 	}
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths II.
// Memory Usage: 2.7 MB, less than 26.53% of Go online submissions for Unique Paths II.

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])

	// 这个矩阵其实就是在原来的grid上面和左边套了一层
	dp := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, col+1)
	}

	// 只是一个初始化
	dp[0][1] = 1

	// 这是按照dp的索引
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			// 这里是将索引映射到原来的grid
			if obstacleGrid[i-1][j-1] != 1 {
				// 如果当前位置没有障碍物，那么就是上面和左边的和
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row][col]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

}
