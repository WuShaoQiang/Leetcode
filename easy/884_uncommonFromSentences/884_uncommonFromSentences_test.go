package uncommonFromSentences

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUncommonFromSentences(t *testing.T) {
	type Input struct {
		A string
		B string
	}
	tests := map[string]struct {
		input Input
		want  []string
	}{
		"Test 1": {
			input: Input{A: "this apple is sweet", B: "this apple is sour"},
			want:  []string{"sweet", "sour"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := uncommonFromSentences(tc.input.A, tc.input.B)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
