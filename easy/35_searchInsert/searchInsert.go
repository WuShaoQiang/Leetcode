package searchInsert

func searchInsert(nums []int, target int) int {
	isFound, idx := binarySearch(nums, 0, len(nums)-1, target)
	if isFound {
		return idx
	}

	if nums[idx] > target {
		return idx
	} else {
		return idx + 1
	}
}

func binarySearch(nums []int, left, right, target int) (found bool, idx int) {
	if left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return true, mid
		} else if nums[mid] > target {
			return binarySearch(nums, left, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	} else if nums[left] == target {
		return true, left
	} else {
		return false, left
	}
}
