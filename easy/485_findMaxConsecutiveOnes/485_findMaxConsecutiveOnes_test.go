package findMaxConsecutiveOnes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 1, 0, 1, 1, 1},
			want:  3,
		},

		"Test 2": {
			input: []int{1, 0, 1, 0, 0, 1},
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
