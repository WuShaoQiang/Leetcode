package numJewelsInStones

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumJewelsInStones(t *testing.T) {
	type Input struct {
		J string
		S string
	}
	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{J: "aA", S: "aAAbbbb"},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numJewelsInStones(tc.input.J, tc.input.S)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
