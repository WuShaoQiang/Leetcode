package findAnagrams

/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

*/

// Runtime: 132 ms, faster than 26.34% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 8 MB, less than 79.61% of Go online submissions for Find All Anagrams in a String.

// func findAnagrams(s string, p string) []int {
// 	if len(s) < len(p) {
// 		return nil
// 	}
// 	pMap := make(map[byte]int)
// 	sMap := make(map[byte]int)

// 	for i := 0; i < len(p); i++ {
// 		pMap[p[i]]++
// 		sMap[s[i]]++
// 	}

// 	res := make([]int, 0)

// 	for i := 0; i < len(s)-len(p); i++ {
// 		if mapEqual(pMap, sMap) {
// 			res = append(res, i)
// 		}
// 		sMap[s[i]]--
// 		sMap[s[i+len(p)]]++
// 	}

// 	if mapEqual(pMap, sMap) {
// 		res = append(res, len(s)-len(p))
// 	}

// 	return res
// }

// func mapEqual(m1, m2 map[byte]int) bool {
// 	for k, v := range m1 {
// 		if v2 := m2[k]; v != v2 {
// 			return false
// 		}
// 	}
// 	return true
// }

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	counts := make([]int, 26)
	for i := 0; i < len(p); i++ {
		counts[int(p[i]-'a')]++
	}

	for i := 0; i < len(p); i++ {
		counts[int(s[i]-'a')]--
	}

	ans := []int{}
	if check(counts) {
		ans = append(ans, 0)
	}

	for i := 1; i+len(p) <= len(s); i++ {
		end := i + len(p) - 1
		counts[int(s[end]-'a')]--
		counts[int(s[i-1]-'a')]++
		if check(counts) {
			ans = append(ans, i)
		}
	}

	return ans
}

func check(counts []int) bool {
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
