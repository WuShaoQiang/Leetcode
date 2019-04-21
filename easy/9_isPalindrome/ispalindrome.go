package isPalindrome

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	l := len(str)
	for i := 0; i < l-1; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}
