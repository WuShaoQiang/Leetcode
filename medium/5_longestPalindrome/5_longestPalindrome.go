package longestPalindrome

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"

*/

// Runtime: 4 ms, faster than 92.64% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.2 MB, less than 72.75% of Go online submissions for Longest Palindromic Substring.

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	max := 1
	start := 0

	for i := 0; i < len(s)-1; i++ {
		expand(s, i, i, &max, &start)
		expand(s, i, i+1, &max, &start)
	}
	return s[start : start+max]
}

func expand(s string, left, right int, max, start *int) {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	dis := right - left - 1
	if dis > *max {
		*max = dis
		*start = left + 1
	}
}
