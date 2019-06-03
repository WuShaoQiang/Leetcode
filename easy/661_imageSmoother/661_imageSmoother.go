package imageSmoother

/*
Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

Example 1:

Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0

Note:

    The value in the given matrix is in the range of [0, 255].
    The length and width of the given matrix are in the range of [1, 150].

*/

// Runtime: 64 ms, faster than 97.96% of Go online submissions for Image Smoother.
// Memory Usage: 7.5 MB, less than 94.12% of Go online submissions for Image Smoother.

func imageSmoother(M [][]int) [][]int {
	r := len(M)
	c := len(M[0])

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum := 0
			count := 0
			for k := j - 1; k < c && k <= j+1; k++ {
				if k < 0 {
					continue
				}
				for t := i - 1; t < r && t <= i+1; t++ {
					if t < 0 {
						continue
					}
					sum += M[t][k]
					count++
				}
			}
			res[i][j] = sum / count
		}
	}

	return res
}
