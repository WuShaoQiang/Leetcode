package removeDuplicates

func removeDuplicates(nums []int) int {
	count := 0
	existMap := make(map[int]bool)
	for _, num := range nums {
		if _, exist := existMap[num]; !exist {
			existMap[num] = true
			nums[count] = num
			count++
		}
	}
	return count
}
