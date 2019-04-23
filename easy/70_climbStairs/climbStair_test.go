package climbStairs

import (
	"fmt"
	"testing"
)

var (
	input = []int{1, 2, 3, 4, 5, 6, 7, 8, 35}
)

func TestClimbStairs(t *testing.T) {
	for _, single := range input {
		fmt.Println(climbStairs(single))
	}
}
