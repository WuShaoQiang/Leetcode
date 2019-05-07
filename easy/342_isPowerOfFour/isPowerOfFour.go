package isPowerOfFour

func isPowerOfFour(num int) bool {
	return num > 0 && 6148914691236517205&num == num && num&(num-1) == 0
}
