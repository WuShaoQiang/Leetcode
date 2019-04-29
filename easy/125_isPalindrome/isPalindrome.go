package isPalindrome

import (
	"strings"
	"unicode"
)

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
