package countPrimes

import (
	"fmt"
	"testing"
)

var (
	input = []int{10, 20, 16, 10000000}
)

func TestCountPrimes(t *testing.T) {
	for _, num := range input {
		fmt.Println(countPrimes(num))
	}
}
