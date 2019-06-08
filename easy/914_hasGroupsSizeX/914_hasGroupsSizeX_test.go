package hasGroupsSizeX

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHasGroupsSizeX(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want:  true,
		},

		"Test 2": {
			input: []int{1, 1, 1, 2, 2, 2, 3, 3},
			want:  false,
		},

		"Test 3": {
			input: []int{1},
			want:  false,
		},

		"Test 4": {
			input: []int{1, 1},
			want:  true,
		},

		"Test 5": {
			input: []int{1, 1, 2, 2, 2, 2},
			want:  true,
		},

		"Test 6": {
			input: []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := hasGroupsSizeX(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
