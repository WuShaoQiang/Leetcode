package findNthDigit

import (
	"fmt"
	"testing"
)

var (
	input = []int{3, 11, 20, 21}
)

func TestFindNthDigit(t *testing.T) {
	for _, single := range input {

		fmt.Println(findNthDigit(single))
	}
}
