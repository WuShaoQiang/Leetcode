package islandPerimeter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIslandPerimeter(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  int
	}{
		"Test 1": {
			input: [][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}},
			want:  16,
		},

		"Test 2": {
			input: [][]int{[]int{1, 0}},
			want:  4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := islandPerimeter(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
