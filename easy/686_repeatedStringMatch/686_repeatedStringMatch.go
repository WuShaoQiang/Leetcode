package repeatedStringMatch

import (
	"strings"
)

/*
Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Repeated String Match.
// Memory Usage: 2.4 MB, less than 74.77% of Go online submissions for Repeated String Match.

func repeatedStringMatch(A string, B string) int {
	if len(A) == 0 {
		return -1
	}

	x := (len(B)-1)/len(A) + 1

	repeated := &strings.Builder{}
	for i := 0; i < x; i++ {
		repeated.WriteString(A)
	}

	// repeat x or x + 1 times for A.
	for i := 0; i < 2; i++ {
		if strings.Contains(repeated.String(), B) {
			return x + i
		}
		repeated.WriteString(A)
	}
	return -1
}
