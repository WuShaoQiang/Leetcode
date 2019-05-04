package containDuplicate

func containsDuplicate(nums []int) bool {
	isExist := make(map[int]bool)

	for _, num := range nums {
		if _, exist := isExist[num]; exist {
			return true
		}

		isExist[num] = true
	}
	return false
}
