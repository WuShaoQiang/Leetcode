package isPalindrome

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true

Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Follow up:

Coud you solve it without converting the integer to a string?
*/

// Runtime: 24 ms, faster than 93.46% of Go online submissions for Palindrome Number.
// Memory Usage: 5.2 MB, less than 39.84% of Go online submissions for Palindrome Number.

// func isPalindrome(x int) bool {
// 	str := strconv.Itoa(x)
// 	for i := 0; i < len(str)/2; i++ {
// 		if str[i] != str[len(str)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// Runtime: 4 ms, faster than 99.81% of Go online submissions for Palindrome Number.
// Memory Usage: 5 MB, less than 72.16% of Go online submissions for Palindrome Number.

func isPalindrome(x int) bool {
	var s int
	for X := x; X > 0; X /= 10 {
		y := X % 10
		s = s*10 + y
	}
	return s == x
}
