package reverseWords

/*
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:

Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Note: In the string, each word is separated by single space and there will not be any extra space in the string
*/

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Reverse Words in a String III.
// Memory Usage: 6.1 MB, less than 54.14% of Go online submissions for Reverse Words in a String III.

func reverseWords(s string) string {
	b := make([]byte, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, reverse(s[start:i])...)
			b = append(b, ' ')
			start = i + 1
		}
	}

	if start != len(s) {
		b = append(b, reverse(s[start:])...)
	}
	return string(b)
}

// Runtime: 8 ms, faster than 75.26% of Go online submissions for Reverse Words in a String III.
// Memory Usage: 6 MB, less than 91.08% of Go online submissions for Reverse Words in a String III.

// func reverseWords(s string) string {
// 	strs := strings.Split(s, " ")
// 	for i := 0; i < len(strs); i++ {
// 		strs[i] = string(reverse(strs[i]))
// 	}
// 	return strings.Join(strs, " ")
// }

func reverse(s string) []byte {
	res := make([]byte, len(s))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = s[i]
	}
	return res
}
