package buddyStrings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuddyStrings(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  bool
	}{
		"Test 1": {
			input: []string{"ab", "ba"},
			want:  true,
		},

		"Test 2": {
			input: []string{"aa", "aa"},
			want:  true,
		},

		"Test 3": {
			input: []string{"ab", "ab"},
			want:  false,
		},

		"Test 4": {
			input: []string{"aaaaaabc", "aaaaaacb"},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := buddyStrings(tc.input[0], tc.input[1])
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
