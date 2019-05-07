package moveZeros

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 1, 0, 3, 12}
)

func TestMoveZeros(t *testing.T) {
	moveZeroes(input)
	fmt.Println(input)
}
