package fairCandySwap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFairCandySwap(t *testing.T) {
	type Input struct {
		A []int
		B []int
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Test 1": {
			input: Input{A: []int{1, 1}, B: []int{2, 2}},
			want:  []int{1, 2},
		},

		"Test 2": {
			input: Input{A: []int{1, 2}, B: []int{2, 3}},
			want:  []int{1, 2},
		},

		"Test 3": {
			input: Input{A: []int{2}, B: []int{1, 3}},
			want:  []int{2, 3},
		},

		"Test 4": {
			input: Input{A: []int{1, 2, 5}, B: []int{2, 4}},
			want:  []int{5, 4},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fairCandySwap(tc.input.A, tc.input.B)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
