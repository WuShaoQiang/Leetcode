package largestTimeFromDigits

import (
	"strconv"
	"strings"
)

/*
Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5.  If no valid time can be made, return an empty string.



Example 1:

Input: [1,2,3,4]
Output: "23:41"

Example 2:

Input: [5,5,5,5]
Output: ""



Note:

    A.length == 4
    0 <= A[i] <= 9


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Largest Time for Given Digits.
// Memory Usage: 2 MB, less than 63.64% of Go online submissions for Largest Time for Given Digits.

func largestTimeFromDigits(A []int) string {
	ans := ""
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			for k := 0; k < len(A); k++ {
				if i == j || i == k || j == k {
					continue
				}
				if A[i]*10+A[j] < 24 && A[k]*10+A[6-i-j-k] < 60 {
					h := strconv.Itoa(A[i]) + strconv.Itoa(A[j])
					m := strconv.Itoa(A[k]) + strconv.Itoa(A[6-i-j-k])
					t := h + ":" + m
					if strings.Compare(t, ans) > 0 {
						ans = t
					}
				}
			}
		}
	}
	return ans
}
