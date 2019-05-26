package validPalindrome

/*
 Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:

Input: "aba"
Output: True

Example 2:

Input: "abca"
Output: True
Explanation: You could delete the character 'c'.

Note:

    The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

*/

// Runtime: 36 ms, faster than 6.57% of Go online submissions for Valid Palindrome II.
// Memory Usage: 7.6 MB, less than 9.09% of Go online submissions for Valid Palindrome II.

// 切片的分割占用了大量的时间
// func validPalindrome(s string) bool {
// 	n := len(s) - 1
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[n-i] {
// 			if i == 0 {
// 				return validation(s[1:]) || validation(s[:n])
// 			}
// 			return validation(s[:i]+s[i+1:]) || validation(s[:n-i]+s[n-i+1:])
// 		}
// 	}
// 	return true
// }

// func validation(s string) bool {
// 	n := len(s) - 1
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[n-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// Runtime: 20 ms, faster than 97.08% of Go online submissions for Valid Palindrome II.
// Memory Usage: 7.3 MB, less than 48.18% of Go online submissions for Valid Palindrome II.

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return validation(s, start+1, end) || validation(s, start, end-1)
		}
		start++
		end--
	}
	return true
}

func validation(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
