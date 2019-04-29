package convertToTitle

import (
	"fmt"
	"testing"
)

var (
	input = []int{1, 28, 52, 701, 702, 28}
)

func TestConvertToTitle(t *testing.T) {
	for _, single := range input {
		fmt.Println(convertToTitle(single))
	}
}
