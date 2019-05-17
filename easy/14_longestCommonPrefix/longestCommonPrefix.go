package longestCommonPrefix

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:

All given inputs are in lowercase letters a-z.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.4 MB, less than 48.70% of Go online submissions for Longest Common Prefix.

func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	}

	minPos := 0
	minLen := len(strs[0])
	for i := 0; i < size; i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
			minPos = i
		}
	}

	for j := 0; j < minLen; j++ {
		temp := strs[minPos][j]
		for i := 0; i < size; i++ {
			if i == minPos {
				continue
			}
			if temp != strs[i][j] {
				if j == 0 {
					return ""
				}
				return strs[i][:j]
			}
		}
	}
	return strs[minPos]

}
