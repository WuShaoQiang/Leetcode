package checkPerfectNumber

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckPerfectNumber(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"Test 1": {
			input: 28,
			want:  true,
		},

		"Test 2": {
			input: 1,
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkPerfectNumber(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
