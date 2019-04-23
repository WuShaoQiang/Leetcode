package removeDuplicates

import (
	"fmt"
	"testing"
)

var (
	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
)

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates(input))
}
