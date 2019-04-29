package twoSum

// 左右夹
func twoSum(numbers []int, target int) []int {
	idx1 := 0
	idx2 := len(numbers) - 1

	// for target < numbers[idx2] {
	// 	idx2--
	// }
	for {
		if idx1 == idx2 {
			return nil
		}
		if result := numbers[idx1] + numbers[idx2]; result != target {
			if result < target {
				idx1++
			} else {
				idx2--
			}
		} else {
			break
		}
	}
	return []int{idx1 + 1, idx2 + 1}
}
