package reverseVowels

/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"

Example 2:

Input: "leetcode"
Output: "leotcede"

Note:
The vowels does not include the letter "y".*/

// Runtime: 4 ms, faster than 97.22% of Go online submissions for Reverse Vowels of a String.
// Memory Usage: 4.1 MB, less than 76.55% of Go online submissions for Reverse Vowels of a String.

func reverseVowels(s string) string {
	b := []byte(s)
	// m := map[byte]bool{
	// 	'a': true,
	// 	'e': true,
	// 	'i': true,
	// 	'o': true,
	// 	'u': true,
	// 	'A': true,
	// 	'E': true,
	// 	'I': true,
	// 	'O': true,
	// 	'U': true,
	// }

	i := 0
	j := len(b) - 1
	for {
		for ; i < len(b); i++ {
			if isVowel(b[i]) {
				break
			}
		}

		for ; j >= 0; j-- {
			if isVowel(b[j]) {
				break
			}
		}

		if i >= j {
			break
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func isVowel(b byte) bool {

	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'

}
