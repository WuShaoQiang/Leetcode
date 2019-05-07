package intersection

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{1, 2, 2, 1},
		[]int{2, 2},
		[]int{4, 9, 5},
		[]int{9, 4, 9, 8, 4},
	}
)

func TestIntersection(t *testing.T) {
	for i := 0; i < len(input); i += 2 {
		fmt.Println(intersection(input[i], input[i+1]))
	}
}
