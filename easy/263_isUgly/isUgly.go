package isUgly


func isUgly(num int) bool {
	if num < 1 {
		return false
	}

	for num%5 == 0 {
		num = num / 5
	}

	for num%3 == 0 {
		num = num / 3
	}

	for num%2 == 0 {
		num = num / 2
	}

	return num == 1
}
