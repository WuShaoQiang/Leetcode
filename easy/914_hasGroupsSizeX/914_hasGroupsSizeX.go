package hasGroupsSizeX

import "math"

func hasGroupsSizeX(deck []int) bool {
	min := math.MaxInt64
	m := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		m[deck[i]]++
	}

	for _, v := range m {
		if min > v {
			min = v
		}
	}

	for _, cnt := range m {
		if division(cnt, min) < 2 {
			return false
		}
	}

	return true
}

func division(x, y int) int {
	// x is bigger than y
	for y != 0 {
		y, x = x%y, y
	}
	return x
}
