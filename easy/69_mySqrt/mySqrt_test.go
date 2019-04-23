package mySqrt

import (
	"fmt"
	"testing"
)

var (
	input = []int{4, 8}
)

func TestMySqrt(t *testing.T) {
	for _, single := range input {
		fmt.Println(mySqrt(single))
	}
}
