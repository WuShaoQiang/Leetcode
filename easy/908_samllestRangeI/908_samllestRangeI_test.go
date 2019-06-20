package samllestRangeI

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSamllestRangeI(t *testing.T) {
	type Input struct {
		A []int
		K int
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{A: []int{1}, K: 0},
			want:  0,
		},

		"Test 2": {
			input: Input{A: []int{0, 10}, K: 2},
			want:  6,
		},

		"Test 3": {
			input: Input{A: []int{1, 3, 6}, K: 3},
			want:  0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := smallestRangeI(tc.input.A, tc.input.K)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
