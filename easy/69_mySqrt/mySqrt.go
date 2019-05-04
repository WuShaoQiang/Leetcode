package mySqrt

import "math"

// func mySqrt(x int) int {
// 	return int(math.Sqrt(float64(x)))
// }

func mySqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}

// func mySqrt(x int) int {
// 	if x == 0 || x == 1 {
// 		return x
// 	}

// 	i := x / 2.0
// 	for i*i > x {
// 		i = (i + x/i) / 2
// 	}
// 	return i
// }
