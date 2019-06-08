package addToArrayForm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddToArrayForm(t *testing.T) {
	type Input struct {
		A []int
		K int
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Test 1": {
			input: Input{A: []int{1, 2, 0, 0}, K: 34},
			want:  []int{1, 2, 3, 4},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := addToArrayForm(tc.input.A, tc.input.K)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
