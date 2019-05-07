package SumRange

import (
	"fmt"
	"testing"
)

var (
	input = []int{-2, 0, 3, -5, 2, -1}
)

func TestSumRange(t *testing.T) {
	a := Constructor(input)
	fmt.Println(a.SumRange(0, 0))
	fmt.Println(a.SumRange(0, 1))
	fmt.Println(a.SumRange(1, 2))
	fmt.Println(a.SumRange(0, 2))
	fmt.Println(a.SumRange(2, 5))
	fmt.Println(a.SumRange(0, 0))
	fmt.Println(a.SumRange(0, 5))
	fmt.Println(a.SumRange(5, 5))
}
