package maxDistToClosest

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaxDistToClosest(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 0, 0, 0, 1, 0, 1},
			want:  2,
		},

		"Test 2": {
			input: []int{1, 0, 0, 0},
			want:  3,
		},

		"Test 3": {
			input: []int{0, 0, 0, 1, 0, 0, 0, 1},
			want:  3,
		},

		"Test 4": {
			input: []int{0, 1, 1, 1, 0, 0, 1, 0, 0},
			want:  2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maxDistToClosest(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
