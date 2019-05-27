package findDisapperedNumbers

/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/

//Runtime: 368 ms, faster than 95.90% of Go online submissions for Find All Numbers Disappeared in an Array.
//Memory Usage: 8.3 MB, less than 65.98% of Go online submissions for Find All Numbers Disappeared in an Array.

func findDisappearedNumbers(nums []int) []int {
    for i:=0;i<len(nums);i++{
        if nums[abs(nums[i])-1] > 0{
            nums[abs(nums[i])-1] *= -1
        }
    }
    
    res := make([]int,0)
    for i:=0;i<len(nums);i++{
        if nums[i] > 0{
            res = append(res,i+1)
        }
    }
    return res
}

func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}