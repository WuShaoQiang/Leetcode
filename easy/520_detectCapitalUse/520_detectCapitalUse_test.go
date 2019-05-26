package detectCapitalUse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDetectCapitalUse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "China",
			want:  true,
		},

		"Test 2": {
			input: "CHina",
			want:  false,
		},

		"Test 3": {
			input: "aeafawA",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := detectCapitalUse(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
