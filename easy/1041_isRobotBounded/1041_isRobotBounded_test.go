package isRobotBounded

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsRobotBounded(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "GGLLGG",
			want:  true,
		},

		"Test 2": {
			input: "GL",
			want:  true,
		},

		"Test 3": {
			input: "GGLL",
			want:  true,
		},

		"Test 4": {
			input: "GGL",
			want:  true,
		},

		"Test 5": {
			input: "GLR",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isRobotBounded(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
