package numSpecialEquivGroups

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumSpecialEquivGroups(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"Test 1": {
			input: []string{"aa", "bb", "cc", "dd"},
			want:  4,
		},

		"Test 2": {
			input: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numSpecialEquivGroups(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
