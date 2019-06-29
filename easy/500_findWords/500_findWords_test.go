package findWords

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindWords(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"Test 1": {
			input: []string{"Hello", "Alaska", "Dad", "Peace"},
			want:  []string{"Alaska", "Dad"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findWords(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
