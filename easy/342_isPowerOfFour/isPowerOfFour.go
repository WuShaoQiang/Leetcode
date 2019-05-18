package isPowerOfFour

/*
Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:

Input: 16
Output: true

Example 2:

Input: 5
Output: false

Follow up: Could you solve it without loops/recursion?*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Power of Four.
// Memory Usage: 2.1 MB, less than 95.51% of Go online submissions for Power of Four.

func isPowerOfFour(num int) bool {
	return num > 0 && 6148914691236517205&num == num && num&(num-1) == 0
}
