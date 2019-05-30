package findUnsortedSubarray

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindUnsortedSubarray(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{2, 6, 4, 8, 10, 9, 15},
			want:  5,
		},

		"Test 2": {
			input: []int{1, 2, 3, 4, 5, 6, 7},
			want:  0,
		},

		"Test 3": {
			input: []int{1, 2, 3, 4, 5, 7, 6},
			want:  2,
		},

		"Test 4": {
			input: []int{1, 2, 3, 4, 5, 7, 6, 8, 9},
			want:  2,
		},

		"Test 5": {
			input: []int{1},
			want:  0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findUnsortedSubarray(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
