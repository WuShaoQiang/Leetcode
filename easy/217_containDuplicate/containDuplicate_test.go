package containDuplicate

import (
	"fmt"
	"testing"
)

var (
	input = []int{1, 2, 3, 1}
)

func TestContainDuplicate(t *testing.T) {
	fmt.Println(containsDuplicate(input))
}
