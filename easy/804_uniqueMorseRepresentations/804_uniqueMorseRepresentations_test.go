package uniqueMorseRepresentations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"Test 1": {
			input: []string{"gin", "zen", "gig", "msg"},
			want:  2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := uniqueMorseRepresentations(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
