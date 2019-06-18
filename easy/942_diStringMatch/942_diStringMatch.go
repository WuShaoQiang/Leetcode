package diStringMatch

/*
Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

    If S[i] == "I", then A[i] < A[i+1]
    If S[i] == "D", then A[i] > A[i+1]



Example 1:

Input: "IDID"
Output: [0,4,1,3,2]

Example 2:

Input: "III"
Output: [0,1,2,3]

Example 3:

Input: "DDI"
Output: [3,2,0,1]



Note:

    1 <= S.length <= 10000
    S only contains characters "I" or "D".

*/

// Runtime: 148 ms, faster than 97.85% of Go online submissions for DI String Match.
// Memory Usage: 8.1 MB, less than 26.23% of Go online submissions for DI String Match.

func diStringMatch(S string) []int {
	left := 0
	right := len(S)
	res := make([]int, len(S)+1)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
	}
	res[len(S)] = left
	return res
}
