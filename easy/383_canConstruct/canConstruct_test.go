package canConstruct

import (
	"fmt"
	"testing"
)

var (
	ransom   = []string{"a", "aa", "aa", "abcdeft"}
	magezine = []string{"b", "ab", "aab", "aabbccddeeff"}
)

func TestCanConstruct(t *testing.T) {
	for i := 0; i < len(ransom); i++ {
		fmt.Println(canConstruct(ransom[i], magezine[i]))
	}
}
