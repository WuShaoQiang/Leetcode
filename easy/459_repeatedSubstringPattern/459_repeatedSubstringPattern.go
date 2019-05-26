package repeatedSubstringPattern

/*
Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.



Example 1:

Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.

Example 2:

Input: "aba"
Output: False

Example 3:

Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)


*/

// Runtime: 8 ms, faster than 76.14% of Go online submissions for Repeated Substring Pattern.
// Memory Usage: 5 MB, less than 78.00% of Go online submissions for Repeated Substring Pattern.

func repeatedSubstringPattern(s string) bool {
	for i := len(s) / 2; i >= 1; i-- {
		if len(s)%i == 0 {
			j := 0
			for j = 0; j < len(s)-i && s[j] == s[j+i]; j++ {

			}
			if j == len(s)-i {
				return true
			}
		}
	}
	return false
}

// Runtime: 4 ms, faster than 98.86% of Go online submissions for Repeated Substring Pattern.
// Memory Usage: 5 MB, less than 86.00% of Go online submissions for Repeated Substring Pattern.

// func repeatedSubstringPattern(s string) bool {
// 	for i := len(s) / 2; i >= 1; i-- {
// 		if len(s)%i == 0 {
// 			tmp := s[0:i]
// 			j := i
// 			for ; j < len(s) && tmp == s[j:j+i]; j += i {

// 			}
// 			if j == len(s) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
