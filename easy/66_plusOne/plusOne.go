package plusOne

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.

Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
// Memory Usage: 2.1 MB, less than 62.67% of Go online submissions for Plus One.

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果不是9,就直接加一返回
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// 执行到这证明这个位是9,需要进位,因此这位变为0
		digits[i] = 0
	}

	// 能运行到这里的只有都是9的情况了,所以要多加一位
	digits = append([]int{1}, digits...)

	return digits
}
