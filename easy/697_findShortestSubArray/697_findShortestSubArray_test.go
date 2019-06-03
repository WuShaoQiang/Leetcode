package findShortestSubArray

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindShortestSubArray(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 2, 2, 3, 1},
			want:  2,
		},

		"Test 2": {
			input: []int{1, 2, 2, 3, 1, 4, 2},
			want:  6,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findShortestSubArray(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
