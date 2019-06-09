package uniquePaths

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniquePaths(t *testing.T) {
	type Input struct {
		m int
		n int
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{m: 3, n: 2},
			want:  3,
		},

		"Test 2": {
			input: Input{m: 7, n: 3},
			want:  28,
		},

		"Test 3": {
			input: Input{m: 1, n: 100},
			want:  1,
		},

		"Test 4": {
			input: Input{m: 3, n: 7},
			want:  28,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := uniquePaths(tc.input.m, tc.input.n)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
