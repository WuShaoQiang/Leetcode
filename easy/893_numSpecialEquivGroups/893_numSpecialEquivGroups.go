package numSpecialEquivGroups

/*
You are given an array A of strings.

Two strings S and T are special-equivalent if after any number of moves, S == T.

A move consists of choosing two indices i and j with i % 2 == j % 2, and swapping S[i] with S[j].

Now, a group of special-equivalent strings from A is a non-empty subset S of A such that any string not in S is not special-equivalent with any string in S.

Return the number of groups of special-equivalent strings from A.



Example 1:

Input: ["a","b","c","a","c","c"]
Output: 3
Explanation: 3 groups ["a","a"], ["b"], ["c","c","c"]

Example 2:

Input: ["aa","bb","ab","ba"]
Output: 4
Explanation: 4 groups ["aa"], ["bb"], ["ab"], ["ba"]

Example 3:

Input: ["abc","acb","bac","bca","cab","cba"]
Output: 3
Explanation: 3 groups ["abc","cba"], ["acb","bca"], ["bac","cab"]

Example 4:

Input: ["abcd","cdab","adcb","cbad"]
Output: 1
Explanation: 1 group ["abcd","cdab","adcb","cbad"]



Note:

    1 <= A.length <= 1000
    1 <= A[i].length <= 20
    All A[i] have the same length.
    All A[i] consist of only lowercase letters.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Groups of Special-Equivalent Strings.
// Memory Usage: 3.2 MB, less than 65.00% of Go online submissions for Groups of Special-Equivalent Strings.

//其实不用调换顺序，只要统计奇数的字符和偶数的字符是否分别一致
func numSpecialEquivGroups(A []string) int {
	m := make(map[string]bool)
	for i := 0; i < len(A); i++ {
		odd := make([]byte, 26)
		even := make([]byte, 26)
		str := A[i]
		for j := 0; j < len(str); j++ {
			if j%2 == 0 {
				even[str[j]-'a']++
			} else {
				odd[str[j]-'a']++
			}
		}
		m[string(even)+string(odd)] = true
	}
	return len(m)
}
