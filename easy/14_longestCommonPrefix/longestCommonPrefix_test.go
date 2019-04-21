package longestCommonPrefix

import (
	"fmt"
	"testing"
)

var (
	input = []string{"caa", "", "a", "acb"}
)

func TestLongestCommonPrefix(t *testing.T) {

	fmt.Println(longestCommonPrefix(input))

}
