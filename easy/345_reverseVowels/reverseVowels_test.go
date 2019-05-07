package reverseVowels

import (
	"fmt"
	"testing"
)

var (
	input = []string{"hello", "leetcode"}
)

func TestReverseVowels(t *testing.T) {
	for _, single := range input {
		fmt.Println(reverseVowels(single))
	}
}
