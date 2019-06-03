package findMaxAverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	type Input struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		input Input
		want  float64
	}{
		"Test 1": {
			input: Input{nums: []int{1, 12, -5, -6, 50, 3}, k: 4},
			want:  12.75,
		},

		"Test 2": {
			input: Input{nums: []int{1, 12}, k: 1},
			want:  12,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findMaxAverage(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, got, "Wrong Answer")
		})
	}
}
