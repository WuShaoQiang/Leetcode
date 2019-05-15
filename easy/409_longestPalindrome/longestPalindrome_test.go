package longestPalindrome

import (
	"fmt"
	"testing"
)

var (
	input = []string{"bananas", "aabb", "abccccdd", "abccdb", "bb", "ccc", "ab", "", "b"}
)

func TestLongestPalindrome(t *testing.T) {
	for _, single := range input {
		fmt.Println(longestPalindrome(single))
	}
}
