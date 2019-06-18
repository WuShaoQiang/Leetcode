package binaryGap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBinaryGap(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 22,
			want:  2,
		},

		"Test 2": {
			input: 5,
			want:  2,
		},

		"Test 3": {
			input: 6,
			want:  1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := binaryGap(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
