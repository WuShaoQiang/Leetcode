package matrixReshape

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMatrixReshape(t *testing.T) {
	type Input struct {
		nums [][]int
		r    int
		c    int
	}
	tests := map[string]struct {
		input Input
		want  [][]int
	}{
		"Test 1": {
			input: Input{nums: [][]int{[]int{1, 2}, []int{3, 4}},
				r: 1,
				c: 4},
			want: [][]int{[]int{1, 2, 3, 4}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := matrixReshape(tc.input.nums, tc.input.r, tc.input.c)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
