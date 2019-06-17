package findErrorNums

/*
 The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.

Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

Example 1:

Input: nums = [1,2,2,4]
Output: [2,3]

Note:

    The given array size will in the range [2, 10000].
    The given array's numbers won't have any order.

*/

// Runtime: 24 ms, faster than 96.12% of Go online submissions for Set Mismatch.
// Memory Usage: 6.1 MB, less than 97.22% of Go online submissions for Set Mismatch.

func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	m := make([]int, len(nums))
	for _, num := range nums {
		m[num-1]++
	}
	for idx, cnt := range m {
		if cnt == 0 {
			res[1] = idx + 1
		}
		if cnt == 2 {
			res[0] = idx + 1
		}
	}
	return res
}
