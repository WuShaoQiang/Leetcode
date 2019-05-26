package countBinarySubStrings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountBinarySubStrings(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"Test 1": {
			input: "00110011",
			want:  6,
		},

		"Test 2": {
			input: "10101",
			want:  4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := countBinarySubstrings(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
