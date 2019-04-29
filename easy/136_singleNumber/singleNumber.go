package singleNumber

func singleNumber(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		nums[i-1] = nums[i-1] ^ nums[i]
	}
	return nums[0]
}
