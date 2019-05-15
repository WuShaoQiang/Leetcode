package longestPalindrome

/*
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.

*/

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	res := 0
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}

	val := 0

	for _, val = range m {
		if val/2 > 0 {
			// 忘记考虑"bananas"这样的情况
			res += val - val%2
		}
	}

	if val == len(s) {
		return val
	}

	if res == len(s) {
		return res
	}

	// 这里没考虑到"bb"这类字符串的情况
	return res + 1
}
