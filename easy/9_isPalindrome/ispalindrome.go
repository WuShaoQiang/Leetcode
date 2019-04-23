package isPalindrome

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// func isPalindrome(x int) bool {
// 	var s int
// 	for X := x; X > 0; X /= 10 {
// 		y := X % 10
// 		s = s*10 + y
// 	}
// 	return s == x
// }
