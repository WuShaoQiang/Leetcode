package addBinary

import (
	"fmt"
	"testing"
)

var (
	a = []string{"11", "1010", "100"}
	b = []string{"1", "1011", "110010"}
)

func TestAddBinary(t *testing.T) {
	for i := 0; i < len(a); i++ {
		fmt.Println(addBinary(a[i], b[i]))
	}
}
