package rob

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{2, 7, 9, 3, 1},
		[]int{10, 9, 1, 9, 1, 9},
		[]int{10, 9, 1, 7, 8, 9, 12, 1, 5, 7},
	}
)

func TestRob(t *testing.T) {
	for _, single := range input {
		fmt.Println(rob(single))
	}
}
