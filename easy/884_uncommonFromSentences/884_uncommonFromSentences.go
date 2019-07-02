package uncommonFromSentences

import "strings"

/*

We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.



Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]


Note:

0 <= A.length <= 200
0 <= B.length <= 200
A and B both contain only spaces and lowercase letters.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Uncommon Words from Two Sentences.
// Memory Usage: 2.3 MB, less than 65.52% of Go online submissions for Uncommon Words from Two Sentences.

func uncommonFromSentences(A string, B string) []string {
	Astrs := strings.Split(A, " ")
	Bstrs := strings.Split(B, " ")
	m := make(map[string]int)

	for _, Astr := range Astrs {
		m[Astr]++
	}

	for _, Bstr := range Bstrs {
		m[Bstr]++
	}

	res := []string{}
	for word, count := range m {
		if count == 1 {
			res = append(res, word)
		}
	}
	return res
}
