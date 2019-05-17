package removeDuplicates

/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.

Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

// Runtime: 40 ms, faster than 99.69% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 7.6 MB, less than 78.23% of Go online submissions for Remove Duplicates from Sorted Array.

// 这个做法没有用到有序数组这个条件
// func removeDuplicates(nums []int) int {
// 	count := 0
// 	existMap := make(map[int]bool)
// 	for _, num := range nums {
// 		if _, exist := existMap[num]; !exist {
// 			existMap[num] = true
// 			nums[count] = num
// 			count++
// 		}
// 	}
// 	return count
// }

// Runtime: 44 ms, faster than 98.69% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 7.6 MB, less than 34.22% of Go online submissions for Remove Duplicates from Sorted Array.

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	before := nums[0]
// 	count := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != before {
// 			nums[count] = nums[i]
// 			before = nums[i]
// 			count++
// 		}
// 	}
// 	return count
// }

// Runtime: 44 ms, faster than 98.69% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 7.5 MB, less than 95.48% of Go online submissions for Remove Duplicates from Sorted Array.

func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[count] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
