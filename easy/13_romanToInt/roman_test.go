package romanToInt

import (
	"fmt"
	"testing"
)

var (
	s = []string{"III", "IV", "LVIII", "MCMXCIV"}
)

func TestRomanToInt(t *testing.T) {
	for _, single := range s {
		fmt.Println(romanToInt(single))
	}
}
