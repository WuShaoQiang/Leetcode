package climbStairs

import "math"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	count2 := n / 2
	result := 0
	for ; count2 >= 0; count2-- {
		result += calc(n-count2, count2)
	}
	return result
}

//注意float的运算可能会出现一些.9999或者.00001的情况
func calc(m, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return m
	}
	if n < (m+1)/2 {
		n = m - n
	}
	tmp := 1.0
	for i := m; i > n; i-- {
		tmp = ((tmp * float64(i)) / (float64(i) - float64(n)))
	}
	return int(math.Round(tmp))
}
