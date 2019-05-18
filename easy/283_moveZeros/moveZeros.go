package moveZeros

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:

    You must do this in-place without making a copy of the array.
    Minimize the total number of operations.

*/

// 类似插入排序
// func moveZeroes(nums []int) {
// 	for idx, num := range nums {
// 		if num == 0 {
// 			continue
// 		}
// 		i := 0
// 		for i = idx; i > 0 && nums[i-1] == 0; i-- {
// 			nums[i] = 0
// 		}

// 		nums[i] = num
// 	}
// }

// Runtime: 60 ms, faster than 99.84% of Go online submissions for Move Zeroes.
// Memory Usage: 7.9 MB, less than 12.65% of Go online submissions for Move Zeroes.

// 将非零的直接往前移，剩下的全部写0
func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}

		nums[i] = nums[j]
		i++
	}

	// for ; i < len(nums); i++ {
	// 	nums[i] = 0
	// }
	copy(nums[i:], make([]int, len(nums)-i))
}
