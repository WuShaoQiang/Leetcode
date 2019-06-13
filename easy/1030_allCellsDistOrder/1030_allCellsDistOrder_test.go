package allCellsDistOrder

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAllCellsDistOrder(t *testing.T) {
	type Input struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := map[string]struct {
		input Input
		want  [][]int
	}{
		"Test 1": {
			input: Input{R: 1, C: 2, r0: 0, c0: 0},
			want:  [][]int{[]int{0, 0}, []int{0, 1}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := allCellsDistOrder(tc.input.R, tc.input.C, tc.input.r0, tc.input.c0)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
