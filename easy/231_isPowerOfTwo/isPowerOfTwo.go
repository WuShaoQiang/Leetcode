package isPowerOfTwo

/*
Given an integer, write a function to determine if it is a power of two.

Example 1:

Input: 1
Output: true 
Explanation: 20 = 1

Example 2:

Input: 16
Output: true
Explanation: 24 = 16

Example 3:

Input: 218
Output: false

*/

// Runtime: 4 ms, faster than 89.96% of Go online submissions for Power of Two.
// Memory Usage: 2.2 MB, less than 66.49% of Go online submissions for Power of Two.

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return (n & (n - 1)) == 0
}
