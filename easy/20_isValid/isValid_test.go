package isValid

import (
	"fmt"
	"testing"
)

var (
	input = []string{"()", "()[]{}", "(]", "([)]", "{[]}", ")}{({))[{{[}", "(([]){})"}
)

func TestIsValid(t *testing.T) {
	for _, single := range input {
		fmt.Println(isValid(single))
	}
}
