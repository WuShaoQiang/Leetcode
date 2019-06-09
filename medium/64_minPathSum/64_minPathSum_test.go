package minPathSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinPathSum(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  int
	}{
		"Test 1": {
			input: [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}},
			want:  7,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := minPathSum(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
