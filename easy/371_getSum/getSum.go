package getSum

/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:

Input: a = 1, b = 2
Output: 3

Example 2:

Input: a = -2, b = 3
Output: 1

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Two Integers.
// Memory Usage: 2 MB, less than 17.31% of Go online submissions for Sum of Two Integers.

// 用&来得到是否有进位，并且要右移1位
// 用^来得到不考虑进位的相加
// 然后进位和不考虑进位的结果再次执行同样的操作，直到没有进位
func getSum(a int, b int) int {
	res := 0
	for {
		carry := (a & b) << 1
		sum := a ^ b
		if carry&sum == 0 {
			res = carry ^ sum
			break
		}
		a = carry
		b = sum
	}
	return res
}
