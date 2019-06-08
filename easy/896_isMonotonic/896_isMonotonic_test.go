package isMonotonic

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsMonotonic(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{1, 2, 2, 3},
			want:  true,
		},

		"Test 2": {
			input: []int{6, 5, 5, 4},
			want:  true,
		},

		"Test 3": {
			input: []int{1, 3, 2},
			want:  false,
		},

		"Test 4": {
			input: []int{1, 2, 4, 5},
			want:  true,
		},

		"Test 5": {
			input: []int{1, 1, 1},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isMonotonic(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
