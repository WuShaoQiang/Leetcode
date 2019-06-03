package largeGroupPositions

/*
In a string S of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.



Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.

Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.

Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]



Note:  1 <= S.length <= 1000

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Positions of Large Groups.
// Memory Usage: 6.1 MB, less than 37.50% of Go online submissions for Positions of Large Groups.

func largeGroupPositions(S string) [][]int {
	start := 0
	currVal := S[0]
	var i int
	res := make([][]int, 0)
	for ; i < len(S); i++ {
		if S[i] != currVal {
			if i-start > 2 {
				res = append(res, []int{start, i - 1})
			}
			start = i
			currVal = S[i]
		}
	}
	if i-start > 2 {
		res = append(res, []int{start, i - 1})
	}
	return res
}
