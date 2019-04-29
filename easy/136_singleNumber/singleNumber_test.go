package singleNumber

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{4, 1, 2, 1, 2},
		[]int{2, 2, 1},
	}
)

func TestSingleNumber(t *testing.T) {
	for _, single := range input {
		fmt.Println(singleNumber(single))

	}
}
