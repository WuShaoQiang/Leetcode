package reverseWords

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseWords(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Test 1": {
			input: "Let's take LeetCode contest",
			want:  "s'teL ekat edoCteeL tsetnoc",
		},

		"Test 2": {
			input: "",
			want:  "",
		},

		"Test 3": {
			input: "a b",
			want:  "a b",
		},

		"Test 4": {
			input: "a b ",
			want:  "a b ",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverseWords(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
