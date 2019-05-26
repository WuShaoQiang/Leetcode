package checkRecord

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckRecord(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "PPALLP",
			want:  true,
		},

		"Test 2": {
			input: "PPALLL",
			want:  false,
		},

		"Test 3": {
			input: "LLALL",
			want:  true,
		},

		"Test 4": {
			input: "LLLAALLL",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkRecord(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
