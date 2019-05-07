package addDigits

import (
	"fmt"
	"testing"
)

var (
	input = []int{38, 49, 111111, 0, 99, 18}
)

func TestAddDigits(t *testing.T) {
	for _, single := range input {

		fmt.Println(addDigits(single))
	}
}
