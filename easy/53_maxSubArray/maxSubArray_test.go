package maxSubArray

import (
	"fmt"
	"testing"
)

var (
	input = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
)

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray(input))
}
