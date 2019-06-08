package sortArrayByParity

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortArrayByParity(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"Test 1": {
			input: []int{3, 1, 2, 4},
			want:  []int{2, 4, 1, 3},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortArrayByParity(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
