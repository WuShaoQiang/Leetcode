package wordPattern

import "strings"

/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true

Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false

Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false

Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false

Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
// Memory Usage: 1.9 MB, less than 55.42% of Go online submissions for Word Pattern.

// 要注意是双边映射
func wordPattern(pattern string, str string) bool {
	patterns := strings.Split(pattern, "")
	strs := strings.Split(str, " ")
	if len(patterns) != len(strs) {
		return false
	}
	pattrenMap := make(map[string]string)
	beUsed := make(map[string]bool)
	for i := 0; i < len(patterns); i++ {
		if _, exist := pattrenMap[patterns[i]]; !exist {
			if _, isUsed := beUsed[strs[i]]; isUsed {
				return false
			}
			pattrenMap[patterns[i]] = strs[i]
			beUsed[strs[i]] = true

		} else {
			if pattrenMap[patterns[i]] != strs[i] {
				return false
			}
		}
	}
	return true

}
