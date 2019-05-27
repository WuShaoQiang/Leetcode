package arrayPairSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArrayPairSum(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 4, 3, 2},
			want:  4,
		},

		"Test 2": {
			input: []int{1, 4, 3, 2, 5, 7},
			want:  9,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := arrayPairSum(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
