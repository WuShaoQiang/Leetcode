package minMoves

import "math"

/*
Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

Example:

Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
*/

// Runtime: 36 ms, faster than 92.59% of Go online submissions for Minimum Moves to Equal Array Elements.
// Memory Usage: 6.6 MB, less than 35.71% of Go online submissions for Minimum Moves to Equal Array Elements.

func minMoves(nums []int) int {
	n := len(nums)
	min := math.MaxInt64
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if min > nums[i] {
			min = nums[i]
		}
	}

	return sum - min*n
}
