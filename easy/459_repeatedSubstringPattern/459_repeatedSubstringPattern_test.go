package repeatedSubstringPattern

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "abab",
			want:  true,
		},

		"Test 2": {
			input: "aba",
			want:  false,
		},

		"Test 3": {
			input: "abcabcabc",
			want:  true,
		},

		"Test 4": {
			input: "",
			want:  false,
		},

		"Test 5": {
			input: "acbcb",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := repeatedSubstringPattern(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
