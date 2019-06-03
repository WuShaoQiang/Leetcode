package checkPossibility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{4, 2, 3},
			want:  true,
		},

		"Test 2": {
			input: []int{4, 2, 1},
			want:  false,
		},

		"Test 3": {
			input: []int{1, 2, 2, 3, 3, 3, 4, 3, 4, 5},
			want:  true,
		},

		"Test 4": {
			input: []int{1, 2, 2, 3, 3, 3, 4, 3, 4, 5, 7, 6},
			want:  false,
		},

		"Test 5": {
			input: []int{3, 4, 2, 3},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkPossibility(tc.input)
			assert.Equal(t, tc.want, got, "Wrong Answer")
		})
	}
}
