package twoSum

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for idx, num := range nums {
		myMap[num] = idx
	}

	for idx, num := range nums {
		tmp := target - num
		if pos, exist := myMap[tmp]; exist {
			if idx > pos {
				return []int{pos, idx}
			} else if idx < pos {
				return []int{idx, pos}
			} else {
				continue
			}
		}
	}
	return nil
}
