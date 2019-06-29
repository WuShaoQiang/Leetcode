package numberOfboomerangs

/*
Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:

Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]

*/

// Runtime: 196 ms, faster than 63.93% of Go online submissions for Number of Boomerangs.
// Memory Usage: 7.5 MB, less than 45.45% of Go online submissions for Number of Boomerangs.

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		m := make(map[int]int, len(points))
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dis := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) +
				(points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			m[dis]++
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}
	return res
}
