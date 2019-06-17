package reachNumber

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReachNumber(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 3,
			want:  2,
		},

		"Test 2": {
			input: 2,
			want:  3,
		},

		"Test 3": {
			input: 1,
			want:  1,
		},

		"Test 4": {
			input: 4,
			want:  3,
		},

		"Test 5": {
			input: -2,
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reachNumber(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
