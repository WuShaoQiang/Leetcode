package findLHS

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindLHS(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 3, 2, 2, 5, 2, 3, 7},
			want:  5,
		},

		"Test 2": {
			input: []int{1, 1, 1, 1},
			want:  0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findLHS(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
