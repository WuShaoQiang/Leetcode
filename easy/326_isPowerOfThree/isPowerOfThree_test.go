package isPowerOfThree

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 3, 9, 27, 100}
)

func TestIsPowerOfThree(t *testing.T) {
	for _, single := range input {
		fmt.Println(isPowerOfThree(single))
	}
}
