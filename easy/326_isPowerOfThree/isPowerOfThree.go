package isPowerOfThree

import "math"

// func isPowerOfThree(n int) bool {
// 	return n > 0 && 4052555153018976267%n == 0
// }

func isPowerOfThree(n int) bool {

	// 多少次方
	k := int(math.Log(math.MaxInt32) / math.Log(3))
	// 求得64位int最大的3的幂
	maxPowerOfThree := int(math.Pow(3, float64(k)))

	return (n > 0 && maxPowerOfThree%n == 0)
}
