package rotatedDiits

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotatedDigits(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 10,
			want:  4,
		},

		"Test 2": {
			input: 20,
			want:  9,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := rotatedDigits(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
