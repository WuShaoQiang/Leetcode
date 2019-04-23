package removeElement

import (
	"fmt"
	"testing"
)

var (
	input = []int{3, 2, 2, 3}
)

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement(input, 3))
}
