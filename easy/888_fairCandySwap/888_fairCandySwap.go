package fairCandySwap

/*
Example 1:

Input: A = [1,1], B = [2,2]
Output: [1,2]

Example 2:

Input: A = [1,2], B = [2,3]
Output: [1,2]

Example 3:

Input: A = [2], B = [1,3]
Output: [2,3]

Example 4:

Input: A = [1,2,5], B = [2,4]
Output: [5,4]



Note:

    1 <= A.length <= 10000
    1 <= B.length <= 10000
    1 <= A[i] <= 100000
    1 <= B[i] <= 100000
    It is guaranteed that Alice and Bob have different total amounts of candy.
    It is guaranteed there exists an answer.


*/

// Runtime: 64 ms, faster than 97.85% of Go online submissions for Fair Candy Swap.
// Memory Usage: 7 MB, less than 43.55% of Go online submissions for Fair Candy Swap.

func fairCandySwap(A []int, B []int) []int {
	BobMap := make(map[int]bool)
	AliceSum := 0
	for i := 0; i < len(A); i++ {
		AliceSum += A[i]
	}

	for i := 0; i < len(B); i++ {
		AliceSum -= B[i]
		BobMap[B[i]] = true
	}

	dis := AliceSum / 2

	for i := 0; i < len(A); i++ {
		if BobMap[A[i]-dis] {
			return []int{A[i], A[i] - dis}
		}
	}

	return nil
}
