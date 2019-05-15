package fizzBuzz

import (
	"fmt"
	"testing"
)

var (
	input = []int{15, 20, 25, 30}
)

func TestFizzBuzz(t *testing.T) {
	for _, single := range input {
		fmt.Println(fizzBuzz(single))
	}
}
