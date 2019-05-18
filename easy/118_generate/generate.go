package generate

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
// Memory Usage: 2.3 MB, less than 31.94% of Go online submissions for Pascal's Triangle.

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
