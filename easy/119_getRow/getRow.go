package getRow

/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]

Follow up:

Could you optimize your algorithm to use only O(k) extra space?

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle II.
// Memory Usage: 2 MB, less than 96.04% of Go online submissions for Pascal's Triangle II.

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1

	for i := 0; i < rowIndex; i++ {
		result[i+1] = result[i]
		for j := i; j > 0; j-- {
			result[j] = result[j] + result[j-1]
		}
	}

	return result
}
