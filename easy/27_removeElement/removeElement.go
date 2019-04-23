package removeElement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		} else {
			i++
		}
	}

	if nums[i] == val {
		return j
	}

	return j + 1

}
