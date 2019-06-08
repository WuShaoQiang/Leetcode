package prefixesDivBy5

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrefixesDivBy5(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []bool
	}{
		"Test 1": {
			input: []int{0, 1, 1},
			want:  []bool{true, false, false},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := prefixesDivBy5(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
