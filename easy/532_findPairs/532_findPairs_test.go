package findPairs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindPairs(t *testing.T) {
	type Input struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{nums: []int{3, 1, 4, 1, 5}, k: 2},
			want:  2,
		},

		"Test 2": {
			input: Input{nums: []int{1, 2, 3, 4, 5}, k: 1},
			want:  4,
		},

		"Test 3": {
			input: Input{nums: []int{1, 3, 1, 5, 4}, k: 0},
			want:  1,
		},

		"Test 4": {
			input: Input{nums: []int{1, 1, 1, 1, 1}, k: 0},
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findPairs(tc.input.nums, tc.input.k)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
