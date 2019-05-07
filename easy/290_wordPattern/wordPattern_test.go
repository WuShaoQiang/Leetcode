package wordPattern

import (
	"fmt"
	"testing"
)

var (
	patterns = []string{"abba", "abba", "aaaa", "abba"}
	strs     = []string{"dog cat cat dog", "dog cat cat fish", "dog cat cat dog", "dog dog dog dog"}
)

func TestWordPattern(t *testing.T) {
	for i := 0; i < len(strs); i++ {
		fmt.Println(wordPattern(patterns[i], strs[i]))
	}

}
