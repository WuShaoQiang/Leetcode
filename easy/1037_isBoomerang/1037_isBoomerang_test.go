package isBoomerang

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsBoomerang(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  bool
	}{
		"Test 1": {
			input: [][]int{[]int{1, 1}, []int{2, 3}, []int{3, 2}},
			want:  true,
		},

		"Test 2": {
			input: [][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}},
			want:  false,
		},

		"Test 3": {
			input: [][]int{[]int{1, 1}, []int{2, 1}, []int{5, 5}},
			want:  true,
		},

		"Test 4": {
			input: [][]int{[]int{1, 1}, []int{1, 1}, []int{5, 5}},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isBoomerang(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
