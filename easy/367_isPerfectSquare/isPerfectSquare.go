package isPerfectSquare

/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true

Example 2:

Input: 14
Output: false

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Perfect Square.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Valid Perfect Square.

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	total := 0
	for i := 1; i <= num/2 && total < num; i++ {
		total += 2*i - 1
	}

	return total == num
}
