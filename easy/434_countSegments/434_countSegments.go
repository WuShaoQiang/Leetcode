package countSegments

/*
Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any non-printable characters.

Example:

Input: "Hello, my name is John"
Output: 5
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Segments in a String.
// Memory Usage: 1.9 MB, less than 70.83% of Go online submissions for Number of Segments in a String.

func countSegments(s string) int {
	count := 0
	letterTrigger := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if !letterTrigger {
				count++
				letterTrigger = true
			}
		} else {
			letterTrigger = false
		}
	}
	return count
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Segments in a String.
// Memory Usage: 1.9 MB, less than 70.83% of Go online submissions for Number of Segments in a String.

// func countSegments(s string) int {
// 	res := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
// 			res++
// 		}
// 	}
// 	return res
// }
