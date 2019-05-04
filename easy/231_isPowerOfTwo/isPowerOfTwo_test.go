package isPowerOfTwo

import (
	"fmt"
	"testing"
)

var (
	input = []int{1, 2, 3, 4, 100, 128}
)

func TestIsPowerOfTwo(t *testing.T) {
	for _, num := range input {
		fmt.Println(isPowerOfTwo(num))
	}
}
