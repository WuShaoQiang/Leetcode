package distributeCandies

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDistributeCandies(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 1, 1, 2, 3, 1},
			want:  3,
		},

		"Test 2": {
			input: []int{1, 1, 1, 1, 3, 1},
			want:  2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := distributeCandies(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
