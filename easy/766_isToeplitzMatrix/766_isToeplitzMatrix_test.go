package isToeplitzMatrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsToeplitzMatrix(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  bool
	}{
		"Test 1": {
			input: [][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}},
			want:  true,
		},

		"Test 2": {
			input: [][]int{[]int{1, 2}, []int{2, 2}},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isToeplitzMatrix(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
