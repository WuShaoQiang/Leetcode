package isIsomorphic

import (
	"fmt"
	"testing"
)

var (
	input = []string{"egg", "add", "foo", "bar", "paper", "title", "ab", "aa"}
)

func TestIsIsomorphic(t *testing.T) {
	for i := 0; i < len(input); i += 2 {
		fmt.Println(isIsomorphic(input[i], input[i+1]))
	}
}
