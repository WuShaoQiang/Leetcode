package rob

func rob(nums []int) int {
	if nums == nil {
		return 0
	}
	beforeLast := 0
	last := 0
	curr := 0

	for i := 0; i < len(nums); i++ {
		curr = max(last, nums[i]+beforeLast)
		beforeLast = last
		last = curr
	}
	return curr

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
