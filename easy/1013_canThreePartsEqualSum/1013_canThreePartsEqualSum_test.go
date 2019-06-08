package canThreePartsEqualSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  bool
	}{
		"Test 1": {
			input: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := canThreePartsEqualSum(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
