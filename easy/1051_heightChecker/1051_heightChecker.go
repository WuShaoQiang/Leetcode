package heightChecker

/*
Students are asked to stand in non-decreasing order of heights for an annual photo.

Return the minimum number of students not standing in the right positions.  (This is the number of students that must move in order for all students to be standing in non-decreasing order of height.)



Example 1:

Input: [1,1,4,2,1,3]
Output: 3
Explanation:
Students with heights 4, 3 and the last 1 are not standing in the right positions.



Note:

    1 <= heights.length <= 100
    1 <= heights[i] <= 100

*/

func heightChecker(heights []int) int {
	origin := append([]int{}, heights...)
	quickSort(heights, 0, len(heights)-1)
	res := 0
	for i := 0; i < len(heights); i++ {
		if origin[i] != heights[i] {
			res++
		}
	}
	return res
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	tmp := arr[left]
	low := left
	high := right
	for low < high {
		for high > low {
			if arr[high] < tmp {
				arr[low] = arr[high]
				low++
				break
			}
			high--
		}
		for low < high {
			if arr[low] > tmp {
				arr[high] = arr[low]
				high--
				break
			}
			low++
		}
	}
	arr[low] = tmp
	quickSort(arr, left, low-1)
	quickSort(arr, low+1, right)
}
