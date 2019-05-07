package isPowerOfFour

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 1, 4, 16, 64, 100, 5}
)

func TestIsPowerOfFour(t *testing.T) {
	for _, single := range input {
		fmt.Println(isPowerOfFour(single))
	}
}
