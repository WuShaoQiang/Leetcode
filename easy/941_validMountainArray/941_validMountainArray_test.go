package validMountainArray

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidMountainArray(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{2, 1},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validMountainArray(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
