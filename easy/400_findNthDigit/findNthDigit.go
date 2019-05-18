package findNthDigit

/*
Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...

Note:
n is positive and will fit within the range of a 32-bit signed integer (n < 231).

Example 1:

Input:
3

Output:
3

Example 2:

Input:
11

Output:
0

Explanation:
The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Nth Digit.
// Memory Usage: 1.9 MB, less than 69.57% of Go online submissions for Nth Digit.

func findNthDigit(n int) int {
	tmp := 9
	i := 1
	for {
		if n <= tmp*i {
			break
		}

		n = n - tmp*i
		tmp = tmp * 10
		i++
	}
	tmp = tmp / 9
	num := tmp + (n-1)/i
	pos := (n - 1) % i
	for j := i - 1; j > pos; j-- {
		num = num / 10
	}
	res := num % 10
	return res
}
