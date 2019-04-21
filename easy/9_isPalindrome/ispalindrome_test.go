package isPalindrome

import (
	"fmt"
	"testing"
)

var (
	input = []int{121, -121, 10}
)

func TestIsPalindrome(t *testing.T) {
	for _, single := range input {
		fmt.Println(isPalindrome(single))
	}
}
