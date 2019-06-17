package checkPerfectNumber

import "math"

/*
We define the Perfect Number is a positive integer that is equal to the sum of all its positive divisors except itself.
Now, given an integer n, write a function that returns true when it is a perfect number and false when it is not.

Example:

Input: 28
Output: True
Explanation: 28 = 1 + 2 + 4 + 7 + 14

Note: The input number n will not exceed 100,000,000. (1e8)
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Perfect Number.
// Memory Usage: 2 MB, less than 5.88% of Go online submissions for Perfect Number.

func checkPerfectNumber(num int) bool {
	if num == 0 {
		return false
	}
	sum := 0
	sn := int(math.Sqrt(float64(num)))
	for i := sn; i >= 1; i-- {
		if num%i == 0 {
			if num/i > i {
				sum += num / i
			}
			sum += i
		}
	}
	return sum == num*2
}
