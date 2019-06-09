package numDecodings

import "strconv"

/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).

Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode Ways.
// Memory Usage: 2 MB, less than 84.35% of Go online submissions for Decode Ways.

// func numDecodings(s string) int {
// 	if len(s) == 0 || s[0] == '0' {
// 		return 0
// 	}

// 	r1 := 1
// 	r2 := 1

// 	for i := 1; i < len(s); i++ {
// 		if s[i] == '0' {
// 			r1 = 0
// 		}

// 		if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6' {
// 			r1 = r2 + r1
// 			r2 = r1 - r2
// 		} else {
// 			r2 = r1
// 		}
// 	}

// 	return r1
// }

func numDecodings(s string) int {
	m := make(map[string]int)

	return way(s, m)
}

func way(s string, m map[string]int) int {
	if len(s) == 0 {
		return 1
	}
	// use map to record
	if _, ok := m[s]; ok {
		return m[s]
	}

	// it will be 0 if the first character is '0'
	// because the whole string is illegal
	if s[0] == '0' {
		return 0
	}

	// s.length=1 and s != '0'
	if len(s) == 1 {
		return 1
	}

	w := way(s[1:], m)
	pre, _ := strconv.Atoi(s[0:2])

	if pre <= 26 {
		w += way(s[2:], m)
	}

	m[s] = w
	return w
}
