package reverseStr

/*
 Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.

Example:

Input: s = "abcdefg", k = 2
Output: "bacdfeg"

Restrictions:

    The string consists of lower English letters only.
    Length of the given string and k will in the range [1, 10000]

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse String II.
// Memory Usage: 4 MB, less than 31.34% of Go online submissions for Reverse String II.

// I misunderstood this problem at the first place, so I did it wrong about seven times before I realized that.
func reverseStr(s string, k int) string {
	if k == 0 {
		return s
	}

	if k >= len(s) {
		return string(reverse(s))
	}

	b := make([]byte, 0)
	skip := 2 * k
	var i int
	for i = 0; i < len(s); {
		if k > len(s)-i {
			b = append(b, reverse(s[i:])...)
		} else {

			b = append(b, reverse(s[i:i+k])...)
			for j := i + k; j < len(s) && j < i+skip; j++ {
				b = append(b, s[j])
			}

		}
		i += skip
	}

	return string(b)
}

func reverse(s string) []byte {
	res := make([]byte, len(s))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[j] = s[i]
	}
	return res
}
