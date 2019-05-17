package lengthOfLastWord

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/

// func lengthOfLastWord(s string) int {
// 	if s == "" {
// 		return 0
// 	}

// 	if !strings.Contains(s, " ") {
// 		return len(s)
// 	}

// 	s = strings.TrimSpace(s)

// 	lastSpaceIndex := strings.LastIndex(s, " ")
// 	return len(s) - lastSpaceIndex - 1
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
// Memory Usage: 2.2 MB, less than 67.23% of Go online submissions for Length of Last Word.

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		// 遇到第一个非空格开始加
		if s[i] != ' ' {
			count++
			// 当遇到空格,并且不是结尾的空格
		} else if s[i] == ' ' && count != 0 {
			return count
		}

	}
	return count
}
