package largestPerimeter

/*
Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.



Example 1:

Input: [2,1,2]
Output: 5

Example 2:

Input: [1,2,1]
Output: 0

Example 3:

Input: [3,2,3,4]
Output: 10

Example 4:

Input: [3,6,2,3]
Output: 8



Note:

    3 <= A.length <= 10000
    1 <= A[i] <= 10^6


*/

// Runtime: 28 ms, faster than 99.16% of Go online submissions for Largest Perimeter Triangle.
// Memory Usage: 6.2 MB, less than 41.43% of Go online submissions for Largest Perimeter Triangle.

func largestPerimeter(A []int) int {
	quickSort(A, 0, len(A)-1)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	low := left
	high := right
	tmp := a[left]
	for low < high {
		for high > low {
			if a[high] < tmp {
				break
			}
			high--
		}

		a[low] = a[high]

		for low < high {
			if a[low] > tmp {
				break
			}
			low++
		}

		a[high] = a[low]
	}
	a[low] = tmp
	quickSort(a, left, low-1)
	quickSort(a, low+1, right)
}
