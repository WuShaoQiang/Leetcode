package diStringMatch

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDiStringMatch(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []int
	}{
		"Test 1": {
			input: "IDID",
			want:  []int{0, 4, 1, 3, 2},
		},

		"Test 2": {
			input: "III",
			want:  []int{0, 1, 2, 3},
		},

		"Test 3": {
			input: "DDI",
			want:  []int{3, 2, 0, 1},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := diStringMatch(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
