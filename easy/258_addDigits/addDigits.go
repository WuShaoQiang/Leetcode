package addDigits

/*
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2 
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2. 
             Since 2 has only one digit, return it.

Follow up:
Could you do it without any loop/recursion in O(1) runtime?*/

// 迭代
// func addDigits(num int) int {
// 	if num < 10 {
// 		return num
// 	}
// 	res := 0
// 	for res == 0 {
// 		for num > 0 {
// 			res += num % 10
// 			fmt.Println("res:", res)
// 			num = num / 10
// 		}
// 		if res >= 10 {
// 			num = res
// 			res = 0
// 		}
// 	}
// 	return res
// }

// Runtime: 4 ms, faster than 89.86% of Go online submissions for Add Digits.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Add Digits.

// f(10*x+y) = f(9*x+x+y) 所以当我们取余9的时候，就成了x+y
// 如果x+y也正好是9的倍数，则返回9

func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	if num%9 == 0 {
		return 9
	}

	return num % 9
}
