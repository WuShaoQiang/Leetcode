package firstUniqChar

import (
	"fmt"
	"testing"
)

var (
	input = []string{"leetcode", "loveleetcode"}
)

func TestFirstUniqChar(t *testing.T) {
	for _, single := range input {
		fmt.Println(firstUniqChar(single))
	}
}
