package rotatedDiits

/*
X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation:
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.

Note:

    N  will be in range [1, 10000].


*/

// Runtime: 16 ms, faster than 15.79% of Go online submissions for Rotated Digits.
// Memory Usage: 2.4 MB, less than 42.42% of Go online submissions for Rotated Digits.

// 转换为字符串做
// func rotatedDigits(N int) int {
// 	res := 0
// 	for i := 0; i <= N; i++ {
// 		s := strconv.Itoa(i)
// 		if validation(s) {
// 			res++
// 		}
// 	}
// 	return res
// }

// func validation(s string) bool {
// 	return (strings.Contains(s, "2") || strings.Contains(s, "5") || strings.Contains(s, "6") || strings.Contains(s, "9")) && !(strings.Contains(s, "3") || strings.Contains(s, "4") || strings.Contains(s, "7"))
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotated Digits.
// Memory Usage: 1.9 MB, less than 90.91% of Go online submissions for Rotated Digits.

// 直接用数字来做
func rotatedDigits(N int) int {
	count := 0
	for i := 0; i <= N; i++ {
		rNum := rotateNum(i)
		if rNum != -1 && rNum != i {
			count++
		}
	}
	return count
}

func rotateNum(num int) int {
	res := 0
	shift := 1
	for num > 0 {
		d := num % 10
		rd := rotateMap(d)
		if rd == -1 {
			return -1
		}
		res += rd * shift
		num /= 10
		shift *= 10
	}
	return res
}

func rotateMap(in int) int {
	m := []int{0, 1, 5, -1, -1, 2, 9, -1, 8, 6}
	return m[in]
}
