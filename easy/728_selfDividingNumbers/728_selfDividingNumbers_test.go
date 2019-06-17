package selfDividingNumbers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSelfDividingNumbers(t *testing.T) {
	type Input struct {
		left  int
		right int
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Test 1": {
			input: Input{left: 1, right: 22},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},

		"Test 2": {
			input: Input{left: 10, right: 10},
			want:  []int{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := selfDividingNumbers(tc.input.left, tc.input.right)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
