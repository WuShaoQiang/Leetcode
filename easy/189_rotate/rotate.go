package rotate

func rotate(nums []int, k int) {
	// 小于等于0直接返回
	if k <= 0 {
		return
	}
	// 如果大于长度，取余数部分再传递一次，记得return
	if k > len(nums) {
		rotate(nums, k%len(nums))
		return
	}

	// 和长度一样等于没变
	if k == len(nums) {
		return
	}
	// 如果数组长度不足以支持翻转，直接返回
	if len(nums) == 1 || len(nums) == 0 {
		return
	}

	// 计算出从哪里分隔
	idx := len(nums) - k
	reverse(nums[:idx])
	reverse(nums[idx:])
	reverse(nums)
}

func reverse(nums []int) {
	l := len(nums)
	// 如果长度为1，不用翻转
	if l == 1 {
		return
	}
	n := (l - 1) / 2
	for n >= 0 {
		nums[n], nums[l-1-n] = nums[l-1-n], nums[n]
		n--
	}
}

// func rotate(nums []int, k int) {
// 	l := len(nums)
// 	for i := l - 1; i > l-1-k; i-- {
// 		tmp := nums[l-1]
// 		for j := l - 1; j > 0; j-- {
// 			nums[j] = nums[j-1]
// 		}
// 		nums[0] = tmp
// 	}
// }
