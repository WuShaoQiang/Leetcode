package flipAndInvertImage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  [][]int
	}{
		"Test 1": {
			input: [][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}},
			want:  [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}},
		},

		"Test 2": {
			input: [][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}},
			want:  [][]int{[]int{1, 1, 0, 0}, []int{0, 1, 1, 0}, []int{0, 0, 0, 1}, []int{1, 0, 1, 0}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := flipAndInvertImage(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
