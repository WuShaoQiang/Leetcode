package minCostClimbingStairs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{10, 15, 20},
			want:  15,
		},

		"Test 2": {
			input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want:  6,
		},

		"Test 3": {
			input: []int{1, 100},
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := minCostClimbingStairs(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
