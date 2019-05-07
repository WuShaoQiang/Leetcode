package findNthDigit

func findNthDigit(n int) int {
	tmp := 9
	i := 1
	for {
		if n <= tmp*i {
			break
		}

		n = n - tmp*i
		tmp = tmp * 10
		i++
	}
	tmp = tmp / 9
	num := tmp + (n-1)/i
	pos := (n - 1) % i
	for j := i - 1; j > pos; j-- {
		num = num / 10
	}
	res := num % 10
	return res
}
