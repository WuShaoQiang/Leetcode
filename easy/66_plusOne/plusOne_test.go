package plusOne

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9, 9, 9},
	}
)

func TestPlusOne(t *testing.T) {
	for _, single := range input {
		fmt.Println(plusOne(single))
	}
}
