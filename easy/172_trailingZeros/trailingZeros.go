package trailingZeros

func trailingZeroes(n int) int {
	count := 0
	tmp := 1
	len := 0
	for n/(tmp*5) > 0 {
		len++
		tmp = tmp * 5
	}
	for i := 0; i < len; i++ {
		count += (n / tmp)
		tmp = tmp / 5
	}
	return count
}
