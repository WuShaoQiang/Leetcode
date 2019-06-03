package isOneBitCharacter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsOneBitCharacter(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{1, 0, 0},
			want:  true,
		},

		"Test 2": {
			input: []int{1, 1, 1, 0},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isOneBitCharacter(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
