package thirdMax

import "math"

// 执行用时 : 16 ms, 在Third Maximum Number的Go提交中击败了27.08% 的用户
// 内存消耗 : 4 MB, 在Third Maximum Number的Go提交中击败了23.33% 的用户

func thirdMax(nums []int) int {
	max := make([]int, 3)
	for i := 0; i < 3; i++ {
		max[i] = math.MinInt64
	}
	count := 0
	isExist := make(map[int]bool)
	for _, num := range nums {
		if _, exist := isExist[num]; !exist {
			count++
			isExist[num] = true
			if num > max[0] {
				copy(max[1:3], max[0:2])
				max[0] = num
			} else if num > max[1] {
				max[2] = max[1]
				max[1] = num
			} else if num > max[2] {
				max[2] = num
			}
		}
	}

	if count < 3 {
		return max[0]
	}

	return max[2]
}

// 执行用时 : 8 ms, 在Third Maximum Number的Go提交中击败了91.67% 的用户
// 内存消耗 : 4 MB, 在Third Maximum Number的Go提交中击败了23.33% 的用户

// func thirdMax(nums []int) int {
// 	max := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}
// 	count := 0
// 	isExist := make(map[int]bool)
// 	for _, num := range nums {
// 		if _, exist := isExist[num]; !exist {
// 			count++
// 			isExist[num] = true
// 			if num > max[0] {
// 				copy(max[1:3], max[0:2])
// 				max[0] = num
// 			} else if num > max[1] {
// 				max[2] = max[1]
// 				max[1] = num
// 			} else if num > max[2] {
// 				max[2] = num
// 			}
// 		}
// 	}

// 	if count < 3 {
// 		return max[0]
// 	}

// 	return max[2]
// }
