package countBinarySubStrings

/*
Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

Example 1:

Input: "00110011"
Output: 6
Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

Notice that some of these substrings repeat and are counted the number of times they occur.

Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

Example 2:

Input: "10101"
Output: 4
Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

Note:
s.length will be between 1 and 50,000.
s will only consist of "0" or "1" characters.
*/

// Runtime: 16 ms, faster than 39.29% of Go online submissions for Count Binary Substrings.
// Memory Usage: 7.6 MB, less than 6.52% of Go online submissions for Count Binary Substrings.

// func countBinarySubstrings(s string) int {
// 	strs := make([]string, 0)
// 	i := 1
// 	start := 0
// 	for ; i < len(s); i++ {
// 		if s[i] != s[i-1] {
// 			strs = append(strs, s[start:i])
// 			start = i
// 		}
// 	}
// 	strs = append(strs, s[start:i])
// 	res := 0
// 	for i = 1; i < len(strs); i++ {
// 		res += min(len(strs[i]), len(strs[i-1]))
// 	}
// 	return res
// }

// Runtime: 8 ms, faster than 100.00% of Go online submissions for Count Binary Substrings.
// Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Count Binary Substrings.

func countBinarySubstrings(s string) int {
	i := 1
	start := 0
	lastLength := 0
	res := 0
	for ; i < len(s); i++ {
		if s[i] != s[i-1] {
			if lastLength == 0 {
				lastLength = i - start
			} else {
				res += min(i-start, lastLength)
				lastLength = i - start
			}
			start = i
		}
	}
	res += min(i-start, lastLength)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
