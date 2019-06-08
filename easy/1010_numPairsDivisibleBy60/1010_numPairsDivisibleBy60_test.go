package numPairsDivisibleBy60

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test 1": {
			input: []int{30, 20, 150, 100, 40},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numPairsDivisibleBy60(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
