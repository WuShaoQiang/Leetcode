package findLHS

/*

We define a harmounious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:

Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].


Note: The length of the input array will not exceed 20,000.
*/

// Runtime: 52 ms, faster than 89.77% of Go online submissions for Longest Harmonious Subsequence.
// Memory Usage: 7.6 MB, less than 9.38% of Go online submissions for Longest Harmonious Subsequence.

func findLHS(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	res := 0
	for k, v := range m {
		if v1, ok := m[k+1]; ok {
			res = max(res, v+v1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
