package allCellsDistOrder

/*
We are given a matrix with R rows and C columns has cells with integer coordinates (r, c), where 0 <= r < R and 0 <= c < C.

Additionally, we are given a cell in that matrix with coordinates (r0, c0).

Return the coordinates of all cells in the matrix, sorted by their distance from (r0, c0) from smallest distance to largest distance.  Here, the distance between two cells (r1, c1) and (r2, c2) is the Manhattan distance, |r1 - r2| + |c1 - c2|.  (You may return the answer in any order that satisfies this condition.)



Example 1:

Input: R = 1, C = 2, r0 = 0, c0 = 0
Output: [[0,0],[0,1]]
Explanation: The distances from (r0, c0) to other cells are: [0,1]

Example 2:

Input: R = 2, C = 2, r0 = 0, c0 = 1
Output: [[0,1],[0,0],[1,1],[1,0]]
Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2]
The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.

Example 3:

Input: R = 2, C = 3, r0 = 1, c0 = 2
Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2,2,3]
There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].



Note:

    1 <= R <= 100
    1 <= C <= 100
    0 <= r0 < R
    0 <= c0 < C


*/

// Runtime: 516 ms, faster than 94.97% of Go online submissions for Matrix Cells in Distance Order.
// Memory Usage: 150.1 MB, less than 63.11% of Go online submissions for Matrix Cells in Distance Order.

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	queue := make([][]int, 0)
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	queue = append(queue, []int{r0, c0})
	visited[r0][c0] = true
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	result := make([][]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		result = append(result, cur)
		queue = queue[1:]
		for _, dir := range dirs {
			r := cur[0] + dir[0]
			c := cur[1] + dir[1]
			if r >= 0 && r < R && c >= 0 && c < C && !visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
			}
		}
	}
	return result
}
