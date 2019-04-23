package searchInsert

import (
	"fmt"
	"testing"
)

var (
	nums    = []int{1, 3, 5, 6}
	targets = []int{5, 2, 7, 0}
)

func TestSearchInsert(t *testing.T) {
	for _, target := range targets {
		fmt.Println(searchInsert(nums, target))
	}
}
