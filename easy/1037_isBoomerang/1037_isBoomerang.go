package isBoomerang

/*
A boomerang is a set of 3 points that are all distinct and not in a straight line.

Given a list of three points in the plane, return whether these points are a boomerang.



Example 1:

Input: [[1,1],[2,3],[3,2]]
Output: true

Example 2:

Input: [[1,1],[2,2],[3,3]]
Output: false



Note:

    points.length == 3
    points[i].length == 2
    0 <= points[i][j] <= 100

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Boomerang.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Valid Boomerang.

func isBoomerang(points [][]int) bool {
	point1 := points[0]
	point2 := points[1]
	point3 := points[2]
	eq := equal(point1, point2) || equal(point1, point3) || equal(point2, point3)
	return !eq && !((point1[1]-point2[1])*(point1[0]-point3[0]) == (point1[1]-point3[1])*(point1[0]-point2[0]))
}

func equal(a, b []int) bool {
	for i := 0; i < 2; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
