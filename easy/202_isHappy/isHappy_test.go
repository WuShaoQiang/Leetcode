package isHappy

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 1, 7, 19, 20, 21}
)

func TestIsHappy(t *testing.T) {
	for _, single := range input {
		fmt.Printf("val:%v  result:%v\n", single, isHappy(single))
	}
}
