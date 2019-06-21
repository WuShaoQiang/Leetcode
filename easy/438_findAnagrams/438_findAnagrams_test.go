package findAnagrams

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindAnagrams(t *testing.T) {
	type Input struct {
		s string
		p string
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Test 1": {
			input: Input{s: "cbaebabacd", p: "abc"},
			want:  []int{0, 6},
		},

		"Test 2": {
			input: Input{s: "abab", p: "ab"},
			want:  []int{0, 1, 2},
		},

		"Test 3": {
			input: Input{s: "ababac", p: "ab"},
			want:  []int{0, 1, 2, 3},
		},

		"Test 4": {
			input: Input{s: "abab", p: "abba"},
			want:  []int{0},
		},

		"Test 5": {
			input: Input{s: "af", p: "be"},
			want:  []int{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findAnagrams(tc.input.s, tc.input.p)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
