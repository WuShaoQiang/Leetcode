package judgeCircle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJudgeCircle(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"Test 1": {
			input: "UUDDLLRR",
			want:  true,
		},

		"Test 2": {
			input: "UUDDUULLLLRRDDRR",
			want:  true,
		},

		"Test 3": {
			input: "UDDLLRR",
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := judgeCircle(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
