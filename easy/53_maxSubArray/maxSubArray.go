package maxSubArray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	total := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if total < 0 {
			total = 0
		}
		total += nums[i]
		if max < total {
			max = total
		}

	}
	return max
}

// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	res := nums[0]
// 	for i := 1; i < n; i++ {
// 		nums[i] = max(nums[i], nums[i]+nums[i-1])
// 		fmt.Println(nums[i])
// 		res = max(res, nums[i])
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
