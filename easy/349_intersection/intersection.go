package intersection

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]bool)
	for _, num1 := range nums1 {
		m[num1] = true
	}

	for _, num2 := range nums2 {
		if use, exist := m[num2]; exist && use {
			res = append(res, num2)
			m[num2] = false
		}
	}
	return res
}
