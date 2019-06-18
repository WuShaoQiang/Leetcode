package isRectangleOverlap

/*
A rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) are the coordinates of its bottom-left corner, and (x2, y2) are the coordinates of its top-right corner.

Two rectangles overlap if the area of their intersection is positive.  To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two (axis-aligned) rectangles, return whether they overlap.

Example 1:

Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
Output: true

Example 2:

Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
Output: false

Notes:

    Both rectangles rec1 and rec2 are lists of 4 integers.
    All coordinates in rectangles will be between -10^9 and 10^9.


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Rectangle Overlap.
// Memory Usage: 1.9 MB, less than 65.52% of Go online submissions for Rectangle Overlap.

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(notOverlap(rec1, rec2) || notOverlap(rec2, rec1))
}

// 比较左边三角形的右上角，与右边三角形的左下角
func notOverlap(leftRec []int, rightRec []int) bool {
	return leftRec[2] <= rightRec[0] || leftRec[3] <= rightRec[1]
}
