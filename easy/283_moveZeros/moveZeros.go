package moveZeros

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
