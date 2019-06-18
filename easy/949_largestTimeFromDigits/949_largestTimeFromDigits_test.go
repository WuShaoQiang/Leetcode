package largestTimeFromDigits

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLargestTimeFromDigits(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  string
	}{
		"Test 1": {
			input: []int{1, 2, 3, 4},
			want:  "23:41",
		},

		"Test 2": {
			input: []int{5, 2, 3, 4},
			want:  "23:54",
		},

		"Test 3": {
			input: []int{5, 5, 5, 5},
			want:  "",
		},

		"Test 4": {
			input: []int{0, 0, 0, 0},
			want:  "00:00",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := largestTimeFromDigits(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
