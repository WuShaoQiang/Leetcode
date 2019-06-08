package transpose

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranspose(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  [][]int
	}{
		"Test 1": {
			input: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			want:  [][]int{[]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}},
		},

		"Test 2": {
			input: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			want:  [][]int{[]int{1, 4}, []int{2, 5}, []int{3, 6}},
		},

		"Test 3": {
			input: [][]int{[]int{1}},
			want:  [][]int{[]int{1}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := transpose(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
