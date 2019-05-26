package isLongPressedName

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsLongPressedName(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  bool
	}{
		"Test 1": {
			input: []string{"alex", "aaleex"},
			want:  true,
		},

		"Test 2": {
			input: []string{"saeed", "ssaaedd"},
			want:  false,
		},

		"Test 3": {
			input: []string{"leelee", "lleeelleee"},
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isLongPressedName(tc.input[0], tc.input[1])
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
