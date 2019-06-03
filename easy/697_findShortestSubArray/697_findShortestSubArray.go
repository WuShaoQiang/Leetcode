package findShortestSubArray

/*
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:

Input: [1, 2, 2, 3, 1]
Output: 2
Explanation:
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.

Example 2:

Input: [1,2,2,3,1,4,2]
Output: 6

Note:
nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.
*/

// Runtime: 28 ms, faster than 93.55% of Go online submissions for Degree of an Array.
// Memory Usage: 7.4 MB, less than 11.11% of Go online submissions for Degree of an Array.

func findShortestSubArray(nums []int) int {
	positionMap := make(map[int][]int)
	countMap := make(map[int]int)
	max := 0
	for i := 0; i < len(nums); i++ {
		if _, exist := positionMap[nums[i]]; !exist {
			positionMap[nums[i]] = make([]int, 2, 2)
			positionMap[nums[i]][0] = i
		} else {
			positionMap[nums[i]][1] = i
		}
		countMap[nums[i]]++
		if countMap[nums[i]] > max {
			max = countMap[nums[i]]
		}
	}

	if max == 1 {
		return 1
	}

	res := len(nums)
	for k, v := range countMap {
		if v == max {
			dis := positionMap[k][1] - positionMap[k][0]
			if res > dis+1 {
				res = dis + 1
			}
		}
	}

	return res
}
