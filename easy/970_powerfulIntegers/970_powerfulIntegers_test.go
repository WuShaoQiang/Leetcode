package powerfulIntegers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPowerfulIntegers(t *testing.T) {
	type Input struct {
		x     int
		y     int
		bound int
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{x: 2, y: 3, bound: 10},
			want:  7,
		},

		"Test 2": {
			input: Input{x: 1, y: 1, bound: 2},
			want:  1,
		},

		"Test 3": {
			input: Input{x: 3, y: 5, bound: 15},
			want:  6,
		},

		"Test 4": {
			input: Input{x: 2, y: 1, bound: 10},
			want:  4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := powerfulIntegers(tc.input.x, tc.input.y, tc.input.bound)
			diff := cmp.Diff(len(got), tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
