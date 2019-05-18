package intersection

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:

    Each element in the result must be unique.
    The result can be in any order.
*/

// Runtime: 4 ms, faster than 97.06% of Go online submissions for Intersection of Two Arrays.
// Memory Usage: 4.4 MB, less than 58.94% of Go online submissions for Intersection of Two Arrays.

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
