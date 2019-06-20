package bitwiseComplement

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBitwiseComplement(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 5,
			want:  2,
		},

		"Test 2": {
			input: 7,
			want:  0,
		},

		"Test 3": {
			input: 0,
			want:  1,
		},

		"Test 4": {
			input: 15,
			want:  0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := bitwiseComplement(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
