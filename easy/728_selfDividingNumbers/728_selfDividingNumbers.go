package selfDividingNumbers

/*
 A self-dividing number is a number that is divisible by every digit it contains.

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

Also, a self-dividing number is not allowed to contain the digit zero.

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

Example 1:

Input:
left = 1, right = 22
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]

Note:
The boundaries of each input argument are 1 <= left <= right <= 10000.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Self Dividing Numbers.
// Memory Usage: 2.7 MB, less than 65.05% of Go online submissions for Self Dividing Numbers.

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		self := i
		isDiv := true
		for self > 0 {
			n := self % 10
			if n == 0 || i%n != 0 {
				isDiv = false
				break
			}
			self /= 10
		}
		if isDiv {
			res = append(res, i)
		}
	}
	return res
}
