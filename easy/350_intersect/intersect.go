package intersect

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for _, num1 := range nums1 {
		m[num1]++
	}

	for _, num2 := range nums2 {
		if count, exist := m[num2]; exist && count > 0 {
			res = append(res, num2)
			m[num2]--
		}
	}
	return res
}
