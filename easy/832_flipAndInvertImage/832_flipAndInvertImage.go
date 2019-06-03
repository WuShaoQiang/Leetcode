package flipAndInvertImage

/*
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Example 1:

Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

Example 2:

Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

Notes:

    1 <= A.length = A[0].length <= 20
    0 <= A[i][j] <= 1


*/

func flipAndInvertImage(A [][]int) [][]int {
	r := len(A)
	c := len(A[0])

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := c - 1; j >= 0; j-- {
			res[i][c-j-1] = 1 - A[i][j]
		}
	}

	return res
}

// Runtime: 4 ms, faster than 98.89% of Go online submissions for Flipping an Image.
// Memory Usage: 3.7 MB, less than 21.46% of Go online submissions for Flipping an Image.

// func flipAndInvertImage(A [][]int) [][]int {
// 	for _, r := range A {
// 		for i := 0; i < len(r)/2; i++ {
// 			r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
// 		}

// 		for i := range r {
// 			r[i] = 1 - r[i]
// 		}
// 	}
// 	return A
// }
