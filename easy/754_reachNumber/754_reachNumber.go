package reachNumber

/*
 You are standing at position 0 on an infinite number line. There is a goal at position target.

On each move, you can either go left or right. During the n-th move (starting from 1), you take n steps.

Return the minimum number of steps required to reach the destination.

Example 1:

Input: target = 3
Output: 2
Explanation:
On the first move we step from 0 to 1.
On the second step we step from 1 to 3.

Example 2:

Input: target = 2
Output: 3
Explanation:
On the first move we step from 0 to 1.
On the second move we step  from 1 to -1.
On the third move we step from -1 to 2.

Note:
target will be a non-zero integer in the range [-10^9, 10^9].
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reach a Number.
// Memory Usage: 1.9 MB, less than 80.00% of Go online submissions for Reach a Number.

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	cnt := 0
	i := 0
	for ; cnt < target; cnt += i {
		i++
	}
	diff := cnt - target
	//如果是偶数，直接返回，因为可以revert一个数字为负数
	if diff%2 == 0 {
		return i
	}
	// 如果是奇数，就要看下一个数是奇数还是偶数
	// 如果是奇数(也就是说当前i为偶数)，那么就加上后面这个数，再反转一个数
	if i%2 == 0 {
		return i + 1
	}

	// 如果后面的那个数是偶数，那么就要加两个数，然后反转一个数
	return i + 2
}
