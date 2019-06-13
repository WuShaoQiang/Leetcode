package arrangeCoins

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArrangeCoins(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 5,
			want:  2,
		},

		"Test 2": {
			input: 8,
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := arrangeCoins(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
