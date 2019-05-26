package checkRecord

/*
You are given a string representing an attendance record for a student. The record only contains the following three characters:

    'A' : Absent.
    'L' : Late.
    'P' : Present.

A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

You need to return whether the student could be rewarded according to his attendance record.

Example 1:

Input: "PPALLP"
Output: True

Example 2:

Input: "PPALLL"
Output: False

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Student Attendance Record I.
// Memory Usage: 2 MB, less than 58.97% of Go online submissions for Student Attendance Record I.

func checkRecord(s string) bool {
	A := 0
	L := 1
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			A++
			if A > 1 {
				return false
			}
		} else if s[i] == 'L' {
			if i > 0 && s[i-1] == 'L' {
				L++
				if L > 2 {
					return false
				}
			} else if i > 0 {
				L = 1
			}
		}
	}

	return true
}
