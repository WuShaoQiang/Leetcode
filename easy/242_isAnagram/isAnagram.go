package isAnagram

/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
// Memory Usage: 3 MB, less than 51.42% of Go online submissions for Valid Anagram.

func isAnagram(s string, t string) bool {
	x := make([]int, 26)
	y := make([]int, 26)

	for i := 0; i < len(s); i++ {
		x[int(s[i]-'a')]++
	}

	for i := 0; i < len(t); i++ {
		y[int(t[i]-'a')]++
	}

	for i := 0; i < 26; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
