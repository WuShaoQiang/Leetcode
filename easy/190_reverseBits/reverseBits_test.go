package reverseBits

import (
	"fmt"
	"testing"
)

var (
	input = []uint32{43261596, 4294967293}
)

func TestReverseBits(t *testing.T) {
	for _, single := range input {
		fmt.Println(reverseBits(single))
	}
}
