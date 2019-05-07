package isPerfectSquare

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	total := 0
	for i := 1; i <= num/2 && total < num; i++ {
		total += 2*i - 1
	}

	return total == num
}
