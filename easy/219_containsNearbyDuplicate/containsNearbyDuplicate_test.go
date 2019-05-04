package containsNearbyDuplicate

import (
	"fmt"
	"testing"
)

var (
	input = [][]int{
		[]int{1, 0, 1, 1},
		[]int{1, 2, 3, 1, 2, 3},
		[]int{1, 2, 3, 1},
	}
)

func TestContainsNearbyDuplicate(t *testing.T) {
	for idx, single := range input {
		fmt.Println(containsNearbyDuplicate(single, idx+1))
	}
}
