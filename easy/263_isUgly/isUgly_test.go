package isUgly

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 1, 25, 6, 8, 15, 14}
)

func TestIsUgly(t *testing.T) {
	for _, single := range input {
		fmt.Println(isUgly(single))
	}
}
