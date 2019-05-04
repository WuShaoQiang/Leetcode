package containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for idx, num := range nums {
		if dis, exist := m[num]; exist && idx-dis <= k {
			return true
		}
		m[num] = idx
	}
	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums) && j <= i+k; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
