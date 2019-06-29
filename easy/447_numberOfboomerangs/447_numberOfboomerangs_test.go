package numberOfboomerangs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumberOfboomerangs(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  int
	}{
		"Test 1": {
			input: [][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}},
			want:  2,
		},

		"Test 2": {
			input: [][]int{[]int{2, 0}, []int{2, 0}, []int{2, 0}, []int{2, 0}},
			want:  24,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numberOfBoomerangs(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
