package maxProfit

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{7, 1, 5, 3, 6, 4},
		[]int{7, 6, 4, 3, 1},
	}
)

func TestMaxProfit(t *testing.T) {
	for _, single := range input {
		fmt.Println(maxProfit(single))
	}
}
