package isHappy

// var isRepeat = make(map[int]bool)

func isHappy(n int) bool {
	isRepeat := make(map[int]bool)
	return ishappy(n, isRepeat)
}

func ishappy(n int, isRepeat map[int]bool) bool {
	if n == 1 {
		return true
	}
	isRepeat[n] = true
	total := 0
	for n > 0 {
		tmp := n % 10
		total += tmp * tmp
		n = n / 10
	}
	if isRepeat[total] == false {
		isRepeat[total] = true
		return ishappy(total, isRepeat)
	}
	return false
}
