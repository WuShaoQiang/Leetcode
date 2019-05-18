package majorityElement

/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3

Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

*/

// Runtime: 16 ms, faster than 99.26% of Go online submissions for Majority Element.
// Memory Usage: 5.9 MB, less than 55.88% of Go online submissions for Majority Element.

func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	count, flag := 1, nums[0]

	for i := 1; i < len(nums); i++ {
		if count < 1 {
			count = 1
			flag = nums[i]
			continue
		}

		if flag == nums[i] {
			count++
		} else {
			count--
		}
	}

	return flag
}

// func majorityElement(nums []int) int {
// 	insertSort(nums)
// 	return nums[len(nums)/2]
// }

// func insertSort(nums []int) {
// 	for i := 1; i < len(nums); i++ {
// 		tmp := nums[i]
// 		j := i - 1
// 		for ; j >= 0 && tmp < nums[j]; j-- {
// 			nums[j+1] = nums[j]
// 		}
// 		nums[j+1] = tmp
// 	}
// }

// func majorityElement(nums []int) int {
// 	halfLength := len(nums) / 2

// 	hashMap := make(map[int]int)

// 	for _, num := range nums {
// 		hashMap[num]++
// 		if hashMap[num] > halfLength {
// 			return num
// 		}
// 	}
// 	return -1
// }
