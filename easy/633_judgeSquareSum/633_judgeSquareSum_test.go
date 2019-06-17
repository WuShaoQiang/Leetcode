package judgeSquareSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJudgeSquareSum(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"Test 1": {
			input: 5,
			want:  true,
		},

		"Test 2": {
			input: 41,
			want:  true,
		},

		"Test 3": {
			input: 40,
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := judgeSquareSum(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
