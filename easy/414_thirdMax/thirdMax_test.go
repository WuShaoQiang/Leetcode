package thirdMax

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{3, 2, 1},
		[]int{1, 2},
		[]int{1, 1, 1, 1, 2},
		[]int{2, 2, 3, 1},
		[]int{1, 2, 3, 9, 10, 11},
		[]int{1, 2, -2147483648},
	}
)

func TestThirdMax(t *testing.T) {
	for _, single := range input {
		fmt.Println(thirdMax(single))
	}
}
