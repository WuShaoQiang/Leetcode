package majorityElement

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
