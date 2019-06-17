package minMoves

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinMoves(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 2, 3},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := minMoves(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
