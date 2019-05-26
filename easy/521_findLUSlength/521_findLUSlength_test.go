package findLUSlength

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindLUSlength(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"Test 1": {
			input: []string{"aba", "cdc"},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findLUSlength(tc.input[0], tc.input[1])
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
