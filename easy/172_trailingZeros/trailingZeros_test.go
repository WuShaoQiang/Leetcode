package trailingZeros

import (
	"fmt"
	"testing"
)

var (
	input = []int{10, 20, 30, 40, 50, 125}
)

func TestTrailingZeros(t *testing.T) {
	for _, num := range input {
		fmt.Println(trailingZeroes(num))

	}
}
