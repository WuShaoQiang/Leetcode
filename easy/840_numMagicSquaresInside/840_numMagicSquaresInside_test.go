package numMagicSquaresInside

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumMagicSquaresInside(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  int
	}{
		"Test 1": {
			input: [][]int{[]int{4, 3, 8, 4}, []int{9, 5, 1, 9}, []int{2, 7, 6, 2}},
			want:  1,
		},

		"Test 2": {
			input: [][]int{[]int{4, 2, 8, 4}, []int{9, 5, 1, 9}, []int{2, 7, 6, 2}},
			want:  0,
		},

		"Test 3": {
			input: [][]int{[]int{5, 5, 5}, []int{5, 5, 5}, []int{5, 5, 5}},
			want:  0,
		},

		"Test 4": {
			input: [][]int{[]int{7, 0, 5}, []int{2, 4, 6}, []int{3, 8, 1}},
			want:  0,
		},

		"Test 5": {
			input: [][]int{[]int{3, 10, 2, 3, 4}, []int{4, 5, 6, 8, 1}, []int{8, 8, 1, 6, 8}, []int{1, 3, 5, 7, 1}, []int{9, 4, 9, 2, 9}},
			want:  1,
		},

		"Test 6": {
			input: [][]int{[]int{8, 1, 6}, []int{3, 5, 7}, []int{4, 9, 2}},
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numMagicSquaresInside(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
