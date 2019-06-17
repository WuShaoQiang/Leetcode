package maxCount

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxCount(t *testing.T) {
	type Input struct {
		m   int
		n   int
		ops [][]int
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{m: 3, n: 3, ops: [][]int{[]int{2, 2}, []int{3, 3}}},
			want:  4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maxCount(tc.input.m, tc.input.n, tc.input.ops)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
