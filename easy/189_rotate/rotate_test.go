package rotate

import (
	"fmt"
	"testing"
)

var (
	input  = []int{1, 2, 3, 4, 5, 6, 7}
	input2 = []int{1}
	input3 = []int{1, 2}
)

func TestRotate(t *testing.T) {
	// rotate(input, 3)
	// fmt.Println(input)
	// rotate(input2, 1)
	// fmt.Println(input2)
	// rotate(nil, 0)
	// rotate(input3, 0)
	// fmt.Println(input3)
	// rotate(input3, 2)
	// fmt.Println(input3)
	rotate(input3, 3)
	fmt.Println(input3)
}
