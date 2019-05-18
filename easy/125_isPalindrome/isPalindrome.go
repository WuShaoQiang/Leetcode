package isPalindrome

import (
	"strings"
	"unicode"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:

Input: "race a car"
Output: false

*/

// Runtime: 4 ms, faster than 87.17% of Go online submissions for Valid Palindrome.
// Memory Usage: 5.1 MB, less than 17.42% of Go online submissions for Valid Palindrome.

func isPalindrome(s string) bool {
	s = strings.Join(filter(s), "")
	// fmt.Println(str)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func filter(s string) []string {
	return strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsNumber(r))
	})
}
