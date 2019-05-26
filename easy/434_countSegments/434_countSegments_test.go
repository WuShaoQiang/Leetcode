package countSegments

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountSegments(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"Test 1": {
			input: "Hello, my name is John",
			want:  5,
		},

		"Test 2": {
			input: "love live! mu'sic forever",
			want:  4,
		},

		"Test 3": {
			input: ", , , , a, eaefa",
			want:  6,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := countSegments(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
