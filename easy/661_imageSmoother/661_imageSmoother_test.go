package imageSmoother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageSmoother(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  [][]int
	}{
		"Test 1": {
			input: [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}},
			want:  [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}},
		},

		"Test 2": {
			input: [][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}},
			want:  [][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := imageSmoother(tc.input)
			assert.Equal(t, tc.want, got, "Wrong Answer")
		})
	}
}
