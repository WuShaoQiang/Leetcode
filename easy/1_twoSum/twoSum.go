package twoSum

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

// Runtime: 4 ms, faster than 99.51% of Go online submissions for Two Sum.
// Memory Usage: 3.8 MB, less than 13.65% of Go online submissions for Two Sum.

// func twoSum(nums []int, target int) []int {
// 	myMap := make(map[int]int)
// 	for idx, num := range nums {
// 		myMap[num] = idx
// 	}

// 	for idx, num := range nums {
// 		tmp := target - num
// 		if pos, exist := myMap[tmp]; exist {
// 			if idx > pos {
// 				return []int{pos, idx}
// 			} else if idx < pos {
// 				return []int{idx, pos}
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	return nil
// }

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Two Sum.
// Memory Usage: 3.7 MB, less than 22.56% of Go online submissions for Two Sum.

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if pos, exist := m[tmp]; exist {
			return []int{pos, i}
		}
		m[nums[i]] = i
	}
	return nil
}
