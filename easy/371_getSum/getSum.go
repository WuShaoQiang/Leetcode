package getSum

// 用&来得到是否有进位，并且要右移1位
// 用^来得到不考虑进位的相加
// 然后进位和不考虑进位的结果再次执行同样的操作，直到没有进位
func getSum(a int, b int) int {
	res := 0
	for {
		carry := (a & b) << 1
		sum := a ^ b
		if carry&sum == 0 {
			res = carry ^ sum
			break
		}
		a = carry
		b = sum
	}
	return res
}
