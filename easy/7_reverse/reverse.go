package reverse

import (
	"math"
)

// -2147483648~2147483647
// func reverse(x int) int {
// 	var rev int
// 	for x != 0 {
// 		lastNum := x % 10
// 		x = x / 10
// 		if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && lastNum > 7 {
// 			return 0
// 		}
// 		if rev < math.MinInt32/10 || rev == math.MinInt32/10 && lastNum < -8 {
// 			return 0
// 		}
// 		rev = rev*10 + lastNum
// 	}
// 	return rev
// }

func core_calc(x int) (ret int) {
	if x < 10 {
		ret = x
	} else {
		for x >= 10 {
			temp := x % 10
			ret += temp
			ret *= 10
			x /= 10
			if x < 10 {
				ret += x
			}
		}
	}
	power := int(math.Pow(2, 31))
	if ret < -power || ret > power {
		ret = 0
	}
	return
}

func reverse(x int) int {
	var ret int
	if x < 0 {
		x *= -1
		ret = core_calc(x)
		ret *= -1
	} else {
		ret = core_calc(x)
	}
	return ret
}
