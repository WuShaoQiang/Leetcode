package dominantIndex

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDominantIndex(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{3, 6, 1, 0},
			want:  1,
		},

		"Test 2": {
			input: []int{1, 2, 3, 4, 5},
			want:  -1,
		},

		"Test 3": {
			input: []int{3},
			want:  0,
		},

		"Test 4": {
			input: []int{3, -1},
			want:  0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := dominantIndex(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
