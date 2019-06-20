package samllestRangeI

/*
Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i].

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.



Example 1:

Input: A = [1], K = 0
Output: 0
Explanation: B = [1]

Example 2:

Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]

Example 3:

Input: A = [1,3,6], K = 3
Output: 0
Explanation: B = [3,3,3] or B = [4,4,4]



Note:

    1 <= A.length <= 10000
    0 <= A[i] <= 10000
    0 <= K <= 10000


*/

// Runtime: 12 ms, faster than 96.47% of Go online submissions for Smallest Range I.
// Memory Usage: 5.9 MB, less than 75.00% of Go online submissions for Smallest Range I.

func smallestRangeI(A []int, K int) int {
	max := A[0]
	min := A[0]
	for i := 1; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		} else if min > A[i] {
			min = A[i]
		}
	}
	if max-min <= 2*K {
		return 0
	}
	return max - min - 2*K

}
