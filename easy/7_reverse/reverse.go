package reverse

import (
	"math"
)

// -2147483648~2147483647
func reverse(x int) int {
	var rev int
	for x != 0 {
		lastNum := x % 10
		x = x / 10
		if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && lastNum > 7 {
			return 0
		}
		if rev < math.MinInt32/10 || rev == math.MinInt32/10 && lastNum < -8 {
			return 0
		}
		rev = rev*10 + lastNum
	}
	return rev
}
