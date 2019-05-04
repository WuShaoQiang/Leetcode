package countPrimes

// func countPrimes(n int) int {
// 	if n < 3 {
// 		return 0
// 	}

// 	flag := make([]bool, n)
// 	num := n / 2
// 	for i := 3; i*i < n; i += 2 {
// 		if flag[i] {
// 			continue
// 		}
// 		for j := i * i; j < n; j += 2 * i {
// 			if !flag[j] {
// 				num--
// 				flag[j] = true
// 			}
// 		}
// 	}
// 	return num

// }

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	flag := make([]bool, n)

	total := n - 2
	// nSqrt := mySqrt(n)
	// 只需要找n的根号前的素数
	for i := 2; i*i < n; i++ {
		if flag[i] {
			continue
		}

		for j := 2; i*j < n; j++ {
			if !flag[i*j] {
				flag[j*i] = true
				total--
			}
		}
	}
	return total
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	i := x / 2.0
	for i*i > x {
		i = (i + x/i) / 2
	}
	return i
}

func findPrimes(n int) []int {
	result := make([]int, n+1)
	result[0] = -1
	result[1] = -1
	for i := n; i > 2; i-- {
		for j := i / 2; j > 1; j-- {
			if i%j == 0 {
				result[i] = -1
				break
			}
		}
	}
	return result
}
