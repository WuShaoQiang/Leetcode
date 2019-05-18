package intersect

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]

Note:

    Each element in the result should appear as many times as it shows in both arrays.
    The result can be in any order.

Follow up:

    What if the given array is already sorted? How would you optimize your algorithm?
    What if nums1's size is small compared to nums2's size? Which algorithm is better?
    What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

*/

// Runtime: 4 ms, faster than 97.98% of Go online submissions for Intersection of Two Arrays II.
// Memory Usage: 4.8 MB, less than 49.30% of Go online submissions for Intersection of Two Arrays II.

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
