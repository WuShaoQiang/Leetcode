package largestPerimeter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLargestPerimeter(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{2, 1, 2},
			want:  5,
		},

		"Test 2": {
			input: []int{1, 2, 1},
			want:  0,
		},

		"Test 3": {
			input: []int{3, 2, 3, 4},
			want:  10,
		},

		"Test 4": {
			input: []int{3, 6, 2, 3},
			want:  8,
		},

		"Test 5": {
			input: []int{3, 6, 2, 3},
			want:  8,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := largestPerimeter(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
