package findLengthOfLCIS

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLengthOfLCIS(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{1, 2, 3, 4, 5},
			want:  5,
		},

		"Test 2": {
			input: []int{1, 2, 3, 3, 5},
			want:  3,
		},

		"Test 3": {
			input: []int{1, 1, 1, 1, 1, 1, 1},
			want:  1,
		},

		"Test 4": {
			input: []int{},
			want:  0,
		},

		"Test 5": {
			input: []int{1},
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findLengthOfLCIS(tc.input)
			assert.Equal(t, tc.want, got, "Wrong Answer")
		})
	}
}
