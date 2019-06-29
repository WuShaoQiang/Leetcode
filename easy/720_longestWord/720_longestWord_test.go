package longestWord

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLongestWord(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"Test 1": {
			input: []string{"w", "wo", "wor", "worl", "world"},
			want:  "world",
		},

		"Test 2": {
			input: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			want:  "apple",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := longestWord(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
