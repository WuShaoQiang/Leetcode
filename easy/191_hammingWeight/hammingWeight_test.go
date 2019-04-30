package hammingWeight

import (
	"fmt"
	"testing"
)

var (
	input = []uint32{1, 2, 3, 4, 5, 6}
)

func TestHammingWeight(t *testing.T) {
	for _, single := range input {
		fmt.Println(hammingWeight(single))
	}
}
