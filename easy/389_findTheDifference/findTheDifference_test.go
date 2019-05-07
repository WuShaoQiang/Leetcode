package findTheDifference

import (
	"fmt"
	"testing"
)

var (
	s = []string{"abcd", "abc"}
	r = []string{"abcde", "adcb"}
)

func TestFindTheDifference(t *testing.T) {
	for i := 0; i < len(s); i++ {

		fmt.Println(string(findTheDifference(s[i], r[i])))
	}
}
