package isPalindrome

import (
	"fmt"
	"testing"
)

var (
	input = "A man, a plan, a canal: Panama"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(input))
}
