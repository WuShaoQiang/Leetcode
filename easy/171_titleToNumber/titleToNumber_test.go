package titleToNumber

import (
	"fmt"
	"testing"
)

var (
	input = []string{"A", "AA", "ZY", "AB"}
)

func TestTitleToNumber(t *testing.T) {
	for _, single := range input {
		fmt.Println(titleToNumber(single))
	}
}
