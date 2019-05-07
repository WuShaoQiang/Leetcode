package missingNumber

// 根据没丢失的和减去丢失数组的和，得到的就是丢失的元素
func missingNumber(nums []int) int {
	total := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		total += i
		sum += nums[i]
	}
	total += len(nums)

	return total - sum
}
