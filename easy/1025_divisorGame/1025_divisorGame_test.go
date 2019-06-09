package divisorGame

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDivisorGame(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"Test 1": {
			input: 2,
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := divisorGame(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
