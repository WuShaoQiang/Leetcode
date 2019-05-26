package toLowerCase

/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"

Example 2:

Input: "here"
Output: "here"

Example 3:

Input: "LOVELY"
Output: "lovely"


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for To Lower Case.
// Memory Usage: 1.9 MB, less than 56.74% of Go online submissions for To Lower Case.

func toLowerCase(str string) string {
	dis := byte(32)
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += dis
		}
	}
	return string(b)
}
