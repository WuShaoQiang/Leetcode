package longestWord

import "strings"

/*
Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.
Example 1:
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
Example 2:
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
Note:

All the strings in the input will only contain lowercase letters.
The length of words will be in the range [1, 1000].
The length of words[i] will be in the range [1, 30].

*/

// Runtime: 20 ms, faster than 40.54% of Go online submissions for Longest Word in Dictionary.
// Memory Usage: 6 MB, less than 42.86% of Go online submissions for Longest Word in Dictionary.

func longestWord(words []string) string {
	m := make(map[string]bool)

	for _, word := range words {
		m[word] = true
	}

	res := ""
	for _, word := range words {
		var i = 1
		for ; i < len(word) && m[word[:i]]; i++ {

		}
		if i == len(word) {
			if len(res) < len(word) {
				res = word
			} else if len(res) == len(word) && strings.Compare(res, word) == 1 {
				res = word
			}
		}
	}
	return res
}

// Runtime: 12 ms, faster than 100.00% of Go online submissions for Longest Word in Dictionary.
// Memory Usage: 5.9 MB, less than 92.86% of Go online submissions for Longest Word in Dictionary.

// func longestWord(words []string) string {
// 	sort.Strings(words)
// 	stack := []string{""}
// 	var res string
// 	for i := 0; i < len(words); i++ {
// 		word := words[i]

// 		var last string
// 		for {
// 			last = stack[len(stack)-1]
// 			if len(stack) == 1 {
// 				break
// 			}
// 			if !(len(last)-len(word) > -1 && last != word[:len(word)-1]) {
// 				break
// 			}
// 			stack = stack[:len(stack)-1]
// 		}
// 		if last != word[:len(word)-1] {
// 			continue
// 		}
// 		stack = append(stack, word)
// 		if len(word) > len(res) {
// 			res = word
// 		}
// 	}
// 	return res
// }
