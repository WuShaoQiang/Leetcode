package reverse

import (
	"fmt"
	"testing"
)

var (
	input = []int{123, -123, 120, 123456789123}
)

func TestReverse(t *testing.T) {
	for _, single := range input {
		fmt.Println(reverse(single))
	}
}
