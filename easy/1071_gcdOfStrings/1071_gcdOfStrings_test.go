package gcdOfStrings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGcdOfStrings(t *testing.T) {
	type Input struct {
		str1 string
		str2 string
	}
	tests := map[string]struct {
		input Input
		want  string
	}{
		"Test 1": {
			input: Input{str1: "ABCABC", str2: "ABC"},
			want:  "ABC",
		},

		"Test 2": {
			input: Input{str1: "ABCABCA", str2: "ABC"},
			want:  "",
		},

		"Test 3": {
			input: Input{str1: "ABABABABAB", str2: "ABABAB"},
			want:  "AB",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := gcdOfStrings(tc.input.str1, tc.input.str2)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
