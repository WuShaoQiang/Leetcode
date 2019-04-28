package mySqrt

import "math"

// func mySqrt(x int) int {
// 	return int(math.Sqrt(float64(x)))
// }

func mySqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}
