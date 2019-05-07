package missingNumber

import (
	"fmt"
	"testing"
)

var (
	input = []int{3, 0, 1}
)

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber(input))
}
