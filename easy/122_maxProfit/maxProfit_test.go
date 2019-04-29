package maxProfit

import (
	"fmt"
	"testing"
)

var (
	input = []int{7, 1, 5, 3, 6, 4}
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit(input))
}
