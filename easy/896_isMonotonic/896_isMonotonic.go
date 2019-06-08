package isMonotonic

/*
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.



Example 1:

Input: [1,2,2,3]
Output: true

Example 2:

Input: [6,5,4,4]
Output: true

Example 3:

Input: [1,3,2]
Output: false

Example 4:

Input: [1,2,4,5]
Output: true

Example 5:

Input: [1,1,1]
Output: true



Note:

    1 <= A.length <= 50000
    -100000 <= A[i] <= 100000


*/

// Runtime: 64 ms, faster than 94.63% of Go online submissions for Monotonic Array.
// Memory Usage: 8.4 MB, less than 58.33% of Go online submissions for Monotonic Array.

func isMonotonic(A []int) bool {
	if len(A) < 3 {
		return true
	}

	monotonic := 0
	for i := 0; i < len(A)-1; i++ {
		if monotonic == 0 {
			if A[i] < A[i+1] {
				monotonic = 1
			} else if A[i] > A[i+1] {
				monotonic = -1
			}
		} else if monotonic == 1 {
			if A[i] > A[i+1] {
				return false
			}
		} else {
			if A[i] < A[i+1] {
				return false
			}
		}
	}

	return true
}
