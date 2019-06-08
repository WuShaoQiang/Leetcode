package addToArrayForm

/*
For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.



Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

Example 2:

Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455

Example 3:

Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021

Example 4:

Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000



Note：

    1 <= A.length <= 10000
    0 <= A[i] <= 9
    0 <= K <= 10000
    If A.length > 1, then A[0] != 0


*/

func addToArrayForm(A []int, K int) []int {
	carry, Ks := 0, splitInt(K)
	if len(A) < len(Ks) {
		A, Ks = Ks, A
	}
	for i, la, lk := 0, len(A)-1, len(Ks)-1; i < len(Ks); i++ {
		num := carry + Ks[lk-i] + A[la-i]
		if num >= 10 {
			num -= 10
			carry = 1

		} else {
			carry = 0
		}
		A[la-i] = num
	}

	for i := len(A) - len(Ks) - 1; i >= 0 && carry == 1; i-- {
		num := carry + A[i]
		if num >= 10 {
			num -= 10
			carry = 1

		} else {
			carry = 0
		}
		A[i] = num
	}

	if carry == 1 {
		A = append([]int{carry}, A...)
	}

	return A
}

func splitInt(K int) []int {
	Ks := make([]int, 0, 5)
	for x := 0; K >= 10 || K <= -10; {
		x = K % 10
		K /= 10
		Ks = append([]int{x}, Ks...)
	}
	Ks = append([]int{K}, Ks...)
	return Ks
}
