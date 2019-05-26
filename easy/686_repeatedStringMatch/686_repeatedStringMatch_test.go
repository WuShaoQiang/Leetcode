package repeatedStringMatch

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepeatedStringMatch(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"Test 1": {
			input: []string{"abcd", "cdabcdab"},
			want:  3,
		},

		"Test 2": {
			input: []string{"abcde", "cdabcdab"},
			want:  -1,
		},

		"Test 3": {
			input: []string{"abcd", "bcdabcdabc"},
			want:  3,
		},

		"Test 4": {
			input: []string{"a", "aa"},
			want:  2,
		},

		"Test 5": {
			input: []string{"abc", "cabcabca"},
			want:  4,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := repeatedStringMatch(tc.input[0], tc.input[1])
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
