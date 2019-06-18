package isRectangleOverlap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsRectangleOverlap(t *testing.T) {
	type Input struct {
		rec1 []int
		rec2 []int
	}
	tests := map[string]struct {
		input Input
		want  bool
	}{
		"Test 1": {
			input: Input{rec1: []int{0, 0, 2, 2}, rec2: []int{1, 1, 3, 3}},
			want:  true,
		},

		"Test 2": {
			input: Input{rec1: []int{0, 0, 1, 1}, rec2: []int{1, 0, 2, 1}},
			want:  false,
		},

		"Test 3": {
			input: Input{rec1: []int{7, 8, 13, 15}, rec2: []int{10, 8, 12, 20}},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isRectangleOverlap(tc.input.rec1, tc.input.rec2)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
