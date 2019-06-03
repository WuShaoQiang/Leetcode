package pivotIndex

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPivotIndex(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 7, 3, 6, 5, 6},
			want:  3,
		},

		"Test 2": {
			input: []int{1, 2, 3},
			want:  -1,
		},

		"Test 3": {
			input: []int{1, 1, 1, 1, 7, 3, 6, 5, 6, 2, 1},
			want:  6,
		},

		"Test 4": {
			input: []int{-1, -1, -1, -1, -1, 0},
			want:  2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := pivotIndex(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
