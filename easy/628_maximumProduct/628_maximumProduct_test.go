package maximumProduct

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaximumProduct(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 2, 3},
			want:  6,
		},

		"Test 2": {
			input: []int{1, 2, 3, 4},
			want:  24,
		},

		"Test 3": {
			input: []int{1, 2, 3, 4, 5, 6, 7},
			want:  210,
		},

		"Test 4": {
			input: []int{-1, -2, -3, 4},
			want:  24,
		},

		"Test 5": {
			input: []int{-1, 2, 3, 4},
			want:  24,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maximumProduct(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
