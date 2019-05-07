package getSum

import (
	"fmt"
	"testing"
)

var (
	a = []int{1, -2, -3}
	b = []int{2, 3, -3}
)

func TestGetSum(t *testing.T) {
	for i := 0; i < len(a); i++ {
		fmt.Println(getSum(a[i], b[i]))
	}
}
