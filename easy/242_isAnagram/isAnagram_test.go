package isAnagram

import (
	"fmt"
	"testing"
)

var (
	s = []string{"anagram", "rat"}
	o = []string{"nagaram", "car"}
)

func TestIsAnagram(t *testing.T) {
	for i := 0; i < len(s); i++ {
		fmt.Println(isAnagram(s[i], o[i]))
	}
}
