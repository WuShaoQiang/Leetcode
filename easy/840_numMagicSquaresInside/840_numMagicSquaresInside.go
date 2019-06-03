package numMagicSquaresInside

/*
A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

Given an grid of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).



Example 1:

Input: [[4,3,8,4],
        [9,5,1,9],
        [2,7,6,2]]
Output: 1
Explanation:
The following subgrid is a 3 x 3 magic square:
438
951
276

while this one is not:
384
519
762

In total, there is only one magic square inside the given grid.

Note:

    1 <= grid.length <= 10
    1 <= grid[0].length <= 10
    0 <= grid[i][j] <= 15


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Magic Squares In Grid.
// Memory Usage: 2.1 MB, less than 55.56% of Go online submissions for Magic Squares In Grid.

func numMagicSquaresInside(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	if r < 3 || c < 3 {
		return 0
	}
	res := 0

	for i := 0; i <= r-3; i++ {
		for j := 0; j <= c-3; j++ {
			targetVal := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			count := 0
			// 横向三次
			exist := make(map[int]bool)
			for ii := i; ii < i+3; ii++ {
				flag := false
				currVal := 0
				for jj := j; jj < j+3; jj++ {
					if !LessThanTen(grid[ii][jj]) || exist[grid[ii][jj]] {
						flag = true
						break
					}
					exist[grid[ii][jj]] = true
					currVal += grid[ii][jj]
				}
				if currVal != targetVal || flag {
					break
				}
				count++
			}
			if count != 3 {
				continue
			}

			// 竖向三次
			for jj := j; jj < j+3; jj++ {
				flag := false
				currVal := 0
				for ii := i; ii < i+3; ii++ {
					if !LessThanTen(grid[ii][jj]) {
						flag = true
						break
					}
					currVal += grid[ii][jj]
				}
				if currVal != targetVal || flag {
					break
				}
				count++
			}
			if count != 6 {
				continue
			}

			// 左对角线一次
			currVal := 0
			for ii, jj := i, j; ii < i+3 && jj < j+3; {
				currVal += grid[ii][jj]
				ii++
				jj++
			}
			if currVal != targetVal {
				continue
			}
			count++

			// 右对角线一次
			currVal = 0
			for ii, jj := i+2, j+2; ii >= i && jj >= j; {
				currVal += grid[ii][jj]
				ii--
				jj--
			}
			if currVal != targetVal {
				continue
			}
			count++
			res++
		}
	}
	return res
}

func LessThanTen(num int) bool {
	return num <= 9 && num > 0
}
