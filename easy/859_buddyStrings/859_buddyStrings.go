package buddyStrings

/*
Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.



Example 1:

Input: A = "ab", B = "ba"
Output: true

Example 2:

Input: A = "ab", B = "ab"
Output: false

Example 3:

Input: A = "aa", B = "aa"
Output: true

Example 4:

Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true

Example 5:

Input: A = "", B = "aa"
Output: false



Note:

    0 <= A.length <= 20000
    0 <= B.length <= 20000
    A and B consist only of lowercase letters.


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Buddy Strings.
// Memory Usage: 2.6 MB, less than 9.52% of Go online submissions for Buddy Strings.

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	// check if it has two same letters
	if A == B {
		m := make(map[byte]bool)
		for i := 0; i < len(A); i++ {
			if m[A[i]] {
				// two times
				return true
			}
			m[A[i]] = true
		}
	}

	start := 0
	end := 0
	// find the two different letters' index
	for {
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				start = i
				break
			}
		}
		for i := start + 1; i < len(A); i++ {
			if A[i] != B[i] {
				end = i
				break
			}
		}

		// if only one different letter
		if start >= end {
			return false
		}

		break
	}

	b := []byte(A)
	b[start], b[end] = b[end], b[start]
	return string(b) == B
}
