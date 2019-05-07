package reverseString

import (
	"fmt"
	"testing"
)

var (
	input = [][]byte{
		[]byte(""),
		[]byte("hello"),
		[]byte("Hannah"),
	}
)

func TestReverseString(t *testing.T) {
	for _, single := range input {
		reverseString(single)
		fmt.Println(string(single))
	}
}
