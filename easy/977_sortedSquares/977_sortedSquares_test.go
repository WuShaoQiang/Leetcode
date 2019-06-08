package sortedSquares

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortedSquares(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"Test 1": {
			input: []int{-4, -1, 0, 3, 10},
			want:  []int{0, 1, 9, 16, 100},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortedSquares(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
