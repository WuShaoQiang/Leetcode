package toLowerCase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToLowerCase(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Test 1": {
			input: "LOVER",
			want:  "lover",
		},

		"Test 2": {
			input: "Lover",
			want:  "lover",
		},

		"Test 3": {
			input: "L$over",
			want:  "l$over",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := toLowerCase(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
