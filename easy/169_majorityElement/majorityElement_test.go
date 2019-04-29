package majorityElement

import (
	"fmt"
	"testing"
)

var (
	input = []int{2, 2, 1, 1, 1, 2, 2}
)

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement(input))
}
