package findErrorNums

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindErrorNums(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"Test 1": {
			input: []int{1, 2, 2, 4},
			want:  []int{2, 3},
		},

		"Test 2": {
			input: []int{1, 2, 3, 3, 4},
			want:  []int{3, 5},
		},

		"Test 3": {
			input: []int{1, 1},
			want:  []int{1, 2},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findErrorNums(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
