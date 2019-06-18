package largestTriangleArea

/*
You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

Example:
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2
Explanation:
The five points are show in the figure below. The red triangle is the largest.

Notes:

    3 <= points.length <= 50.
    No points will be duplicated.
     -50 <= points[i][j] <= 50.
    Answers within 10^-6 of the true value will be accepted as correct.

*/

// Runtime: 4 ms, faster than 35.29% of Go online submissions for Largest Triangle Area.
// Memory Usage: 2.2 MB, less than 33.33% of Go online submissions for Largest Triangle Area.

func largestTriangleArea(points [][]int) float64 {
	max := float64(0)
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := i + 2; k < len(points); k++ {
				if tmp := calc(points[i][0], points[j][0], points[k][0], points[i][1], points[j][1], points[k][1]); max < tmp {
					max = tmp
				}
			}
		}
	}
	return max
}

func calc(x1, x2, x3, y1, y2, y3 int) float64 {
	return abs(float64(x1*y2-x2*y1+x2*y3-x3*y2+x3*y1-x1*y3) / 2)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
