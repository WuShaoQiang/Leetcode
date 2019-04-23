package twoSum

import (
	"fmt"
	"testing"
)

var (
	input = []int{2, 7, 11, 15}
)

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum(input, 9))
}
