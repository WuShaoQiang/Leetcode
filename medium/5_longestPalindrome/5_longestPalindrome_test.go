package longestPalindrome

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLongestPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Test 1": {
			input: "babad",
			want:  "bab",
		},

		"Test 2": {
			input: "cbbd",
			want:  "bb",
		},

		"Test 3": {
			input: "ac",
			want:  "a",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := longestPalindrome(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
