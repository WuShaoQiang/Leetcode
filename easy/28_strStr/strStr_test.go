package strStr

import (
	"fmt"
	"testing"
)

var (
	haystack = []string{"hello", "aaaaa", "mississippi"}
	needle   = []string{"ll", "bba", "sipp"}
)

func TestStrStr(t *testing.T) {
	for i := 0; i < len(haystack); i++ {
		fmt.Println(strStr(haystack[i], needle[i]))
	}
}
