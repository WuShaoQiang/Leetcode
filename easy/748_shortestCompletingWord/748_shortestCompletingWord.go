package shortestCompletingWord

import "strings"

/*
Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate. Such a word is said to complete the given string licensePlate

Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.

It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.

The license plate might have the same letter occurring multiple times. For example, given a licensePlate of "PP", the word "pair" does not complete the licensePlate, but the word "supper" does.

Example 1:
Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
Output: "steps"
Explanation: The smallest length word that contains the letters "S", "P", "S", and "T".
Note that the answer is not "step", because the letter "s" must occur in the word twice.
Also note that we ignored case for the purposes of comparing whether a letter exists in the word.
Example 2:
Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
Output: "pest"
Explanation: There are 3 smallest length words that contains the letters "s".
We return the one that occurred first.
Note:
licensePlate will be a string with length in range [1, 7].
licensePlate will contain digits, spaces, or letters (uppercase or lowercase).
words will have a length in the range [10, 1000].
Every words[i] will consist of lowercase letters, and have length in range [1, 15].
*/

func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0; i < len(licensePlate); i++ {
		b := licensePlate[i]
		if b >= 'a' && b <= 'z' {
			m[b]++
		}
	}
	res := ""
	resLen := 1001
	for _, word := range words {
		word = strings.ToLower(word)
		wordMap := make(map[byte]int)
		for i := 0; i < len(word); i++ {
			wordMap[word[i]]++
		}
		flag := true
		for k, v := range m {
			if v > wordMap[k] {
				flag = false
				break
			}
		}
		if flag {
			if len(word) < resLen {
				res = word
				resLen = len(word)
			}
		}
	}
	return res
}

// Runtime: 4 ms, faster than 100.00% of Go online submissions for Shortest Completing Word.
// Memory Usage: 4 MB, less than 100.00% of Go online submissions for Shortest Completing Word.

// func shortestCompletingWord(licensePlate string, words []string) string {
// 	var res string
// 	counter := countLetters(licensePlate)
// 	for _, word := range words {
// 		if contains(countLetters(word), counter) && (res == "" || len(word) < len(res)) {
// 			res = word
// 		}
// 	}
// 	return res
// }

// func countLetters(s string) [26]int {
// 	var res [26]int
// 	for i := 0; i < len(s); i++ {
// 		c := s[i]
// 		if c >= 'A' && c <= 'Z' {
// 			res[c-'A']++
// 		} else if c >= 'a' && c <= 'z' {
// 			res[c-'a']++
// 		}
// 	}
// 	return res
// }

// func contains(a, b [26]int) bool {
// 	for key, valueB := range b {
// 		valueA := a[key]
// 		if valueA < valueB {
// 			return false
// 		}
// 	}
// 	return true
// }
