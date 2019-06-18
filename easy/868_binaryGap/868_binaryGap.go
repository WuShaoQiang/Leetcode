package binaryGap

/*
Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.



Example 1:

Input: 22
Output: 2
Explanation:
22 in binary is 0b10110.
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
The first consecutive pair of 1's have distance 2.
The second consecutive pair of 1's have distance 1.
The answer is the largest of these two distances, which is 2.

Example 2:

Input: 5
Output: 2
Explanation:
5 in binary is 0b101.

Example 3:

Input: 6
Output: 1
Explanation:
6 in binary is 0b110.

Example 4:

Input: 8
Output: 0
Explanation:
8 in binary is 0b1000.
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.



Note:

    1 <= N <= 10^9


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Gap.
// Memory Usage: 2 MB, less than 65.00% of Go online submissions for Binary Gap.

func binaryGap(N int) int {
	max := 0
	var i = 1
	for i <= 64 && N != 0 {
		start := 0
		end := 0
		for ; i <= 64; i++ {
			if N&1 == 1 {
				start = i
				i++
				N >>= 1
				break
			}
			N >>= 1
		}

		for ; i <= 64; i++ {
			if N&1 == 1 {
				end = i
				break
			}
			N >>= 1
		}

		// fmt.Printf("N %v\n", N)
		// fmt.Printf("start %v end %v\n", start, end)

		if start != 0 && end != 0 {
			if max < end-start {
				max = end - start
			}
		}
	}

	return max
}
