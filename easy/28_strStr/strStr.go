package strStr

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
// Memory Usage: 2.3 MB, less than 49.83% of Go online submissions for Implement strStr().

// func strStr(haystack string, needle string) int {
// 	return strings.Index(haystack, needle)
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
// Memory Usage: 2.3 MB, less than 62.80% of Go online submissions for Implement strStr().

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 检测到第一个字符相等(当初忘记考虑只有一个字符的情况)
		if haystack[i] == needle[0] {
			// 如果needle的字符只有一个,直接返回i
			if len(needle) == 1 {
				return i
			}
			// 从1~len(needle)-1
			for j := 1; j < len(needle) && haystack[i+j] == needle[j]; j++ {
				// 如果最后一个字符是相等的,就找到了
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	// 找不到返回-1
	return -1
}
