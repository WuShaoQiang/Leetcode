package toHex

import (
	"fmt"
	"testing"
)

var (
	input = []int{26, -1, 0, 100,255}
)

func TestToHex(t *testing.T) {
	for _, single := range input {
		fmt.Println(toHex(single))
	}
}
