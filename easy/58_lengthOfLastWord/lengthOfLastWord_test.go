package lengthOfLastWord

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	inpit = "Hello World"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"Test 1": {
			input: "Hello World",
			want:  5,
		},

		"Test 2": {
			input: " ",
			want:  0,
		},

		"Test 3": {
			input: " ",
			want:  0,
		},

		"Test 4": {
			input: "a",
			want:  1,
		},

		"Test 5": {
			input: "a b c",
			want:  1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lengthOfLastWord(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
