package isPerfectSquare

import (
	"fmt"
	"testing"
)

var (
	input = []int{1, 4, 9, 16, 25, 100, 101}
)

func TestIsPerfectSquare(t *testing.T) {
	for _, single := range input {
		fmt.Println(isPerfectSquare(single))
	}
}
