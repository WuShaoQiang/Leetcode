package reverseStr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseStr(t *testing.T) {
	type Input struct {
		s string
		k int
	}
	tests := map[string]struct {
		input Input
		want  string
	}{
		"Test 1": {
			input: Input{s: "abcdefg", k: 2},
			want:  "bacdfeg",
		},

		"Test 3": {
			input: Input{s: "abcdef", k: 3},
			want:  "cbadef",
		},

		"Test 4": {
			input: Input{s: "abcdef", k: 6},
			want:  "fedcba",
		},

		"Test 5": {
			input: Input{s: "a", k: 2},
			want:  "a",
		},

		"Test 6": {
			input: Input{s: "abcdefg", k: 3},
			want:  "cbadefg",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverseStr(tc.input.s, tc.input.k)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
