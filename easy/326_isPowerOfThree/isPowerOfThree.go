package isPowerOfThree

import "math"

/*
Given an integer, write a function to determine if it is a power of three.

Example 1:

Input: 27
Output: true

Example 2:

Input: 0
Output: false

Example 3:

Input: 9
Output: true

Example 4:

Input: 45
Output: false

Follow up:
Could you do it without using any loop / recursion?*/

// func isPowerOfThree(n int) bool {
// 	return n > 0 && 4052555153018976267%n == 0
// }

// Runtime: 20 ms, faster than 98.68% of Go online submissions for Power of Three.
// Memory Usage: 6 MB, less than 96.63% of Go online submissions for Power of Three.

func isPowerOfThree(n int) bool {

	// 多少次方
	k := int(math.Log(math.MaxInt32) / math.Log(3))
	// 求得64位int最大的3的幂
	maxPowerOfThree := int(math.Pow(3, float64(k)))

	return (n > 0 && maxPowerOfThree%n == 0)
}
