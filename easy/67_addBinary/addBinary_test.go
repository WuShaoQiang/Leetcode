package addBinary

import (
	"fmt"
	"testing"
)

var (
	a = []string{"11", "1010"}
	b = []string{"1", "1011"}
)

func TestAddBinary(t *testing.T) {
	for i := 0; i < len(a); i++ {
		fmt.Println(addBinary(a[i], b[i]))
	}
}
