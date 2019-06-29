package islandPerimeter

/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.



Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:
*/

// Runtime: 56 ms, faster than 95.97% of Go online submissions for Island Perimeter.
// Memory Usage: 6.4 MB, less than 75.86% of Go online submissions for Island Perimeter.

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 1 {
				res += islandHelper(grid, x, y)
			}
		}
	}
	return res
}

func islandHelper(grid [][]int, x, y int) int {
	res := 4
	//left
	if x-1 >= 0 && grid[y][x-1] == 1 {
		res--
	}
	//right
	if x+1 < len(grid[0]) && grid[y][x+1] == 1 {
		res--
	}
	//up
	if y-1 >= 0 && grid[y-1][x] == 1 {
		res--
	}
	//down
	if y+1 < len(grid) && grid[y+1][x] == 1 {
		res--
	}

	return res
}
