package validPalindrome

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "aba",
			want:  true,
		},

		"Test 2": {
			input: "abca",
			want:  true,
		},

		"Test 3": {
			input: "abcca",
			want:  true,
		},

		"Test 4": {
			input: "abcda",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validPalindrome(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
