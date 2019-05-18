package SumRange

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:

Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

Note:

    You may assume that the array does not change.
    There are many calls to sumRange function.
*/

// 自己的方法

// type NumArray struct {
// 	arr     []int
// 	lastSum int
// 	lasti   int
// 	lastj   int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{
// 		arr: nums,
// 	}
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	if i == j {
// 		return this.arr[i]
// 	}

// 	currSum := this.lastSum
// 	if this.lasti != i {
// 		if i < this.lasti { //与前面相加
// 			for idx := i; idx < this.lasti; idx++ {
// 				currSum += this.arr[idx]
// 			}
// 		} else { //与前面相减
// 			for idx := this.lasti; idx < i; idx++ {
// 				currSum -= this.arr[idx]
// 			}
// 		}
// 	}

// 	if this.lastj != j {
// 		if j < this.lastj { //与后面相减
// 			for idx := j + 1; idx <= this.lastj; idx++ {
// 				currSum -= this.arr[idx]
// 			}
// 		} else if this.lastj == 0 { //与后面相加
// 			for idx := this.lastj; idx <= j; idx++ {
// 				currSum += this.arr[idx]
// 			}
// 		} else {
// 			for idx := this.lastj + 1; idx <= j; idx++ {
// 				currSum += this.arr[idx]
// 			}
// 		}
// 	}

// 	this.lastSum = currSum
// 	this.lasti = i
// 	this.lastj = j

// 	return currSum
// }

// Runtime: 36 ms, faster than 88.09% of Go online submissions for Range Sum Query - Immutable.
// Memory Usage: 10.7 MB, less than 22.22% of Go online submissions for Range Sum Query - Immutable.

// 参考最优答案
type NumArray struct {
	arr     []int
	sumList []int
	len     int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	if l == 0 {
		return NumArray{}
	}
	sumList := make([]int, l)
	sumList[0] = nums[0]
	for i := 1; i < l; i++ {
		sumList[i] = sumList[i-1] + nums[i]
	}
	return NumArray{
		arr:     nums,
		sumList: sumList,
		len:     l,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sumList[j]
	}

	return this.sumList[j] - this.sumList[i-1]
}
