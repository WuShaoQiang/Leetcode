package sortArrayByParity

/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.



Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.



Note:

    1 <= A.length <= 5000
    0 <= A[i] <= 5000


*/

// Runtime: 52 ms, faster than 96.47% of Go online submissions for Sort Array By Parity.
// Memory Usage: 8 MB, less than 86.45% of Go online submissions for Sort Array By Parity.

func sortArrayByParity(A []int) []int {
	left := 0
	right := len(A) - 1

	res := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[left] = A[i]
			left++
		} else {
			res[right] = A[i]
			right--
		}
	}

	return res
}
