package largestTriangleArea

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLargestTriangleArea(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  float64
	}{
		"Test 1": {
			input: [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{0, 2}, []int{2, 0}},
			want:  float64(2),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := largestTriangleArea(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
