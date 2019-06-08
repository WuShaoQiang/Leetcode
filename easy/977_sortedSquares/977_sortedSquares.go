package sortedSquares

/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.



Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]



Note:

    1 <= A.length <= 10000
    -10000 <= A[i] <= 10000
    A is sorted in non-decreasing order.


*/

func sortedSquares(A []int) []int {
	right := len(A) - 1
	left := 0
	resPointer := len(A) - 1
	res := make([]int, len(A))
	for left < right {
		if abs(A[left]) <= abs(A[right]) {
			res[resPointer] = A[right] * A[right]
			right--
		} else {
			res[resPointer] = A[left] * A[left]
			left++
		}
		resPointer--
	}
	res[0] = A[left] * A[left]
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
