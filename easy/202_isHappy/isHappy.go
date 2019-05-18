package isHappy

/*
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example: 

Input: 19
Output: true
Explanation: 
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
// Memory Usage: 2.1 MB, less than 42.17% of Go online submissions for Happy Number.

func isHappy(n int) bool {
	isRepeat := make(map[int]bool)
	return ishappy(n, isRepeat)
}

func ishappy(n int, isRepeat map[int]bool) bool {
	if n == 1 {
		return true
	}
	isRepeat[n] = true
	total := 0
	for n > 0 {
		tmp := n % 10
		total += tmp * tmp
		n = n / 10
	}
	if isRepeat[total] == false {
		isRepeat[total] = true
		return ishappy(total, isRepeat)
	}
	return false
}
