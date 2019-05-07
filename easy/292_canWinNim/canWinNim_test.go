package canWinNim

import (
	"fmt"
	"testing"
)

var (
	input = []int{4, 6, 8, 10, 12}
)

func TestCanWinNim(t *testing.T) {
	for _, single := range input {

		fmt.Println(canWinNim(single))
	}
}
