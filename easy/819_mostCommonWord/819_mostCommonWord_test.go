package mostCommonWord

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMostCommonWord(t *testing.T) {
	type Input struct {
		paragraph string
		banned    []string
	}
	tests := map[string]struct {
		input Input
		want  string
	}{
		"Test 1": {
			input: Input{paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.", banned: []string{"hit"}},
			want:  "ball",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mostCommonWord(tc.input.paragraph, tc.input.banned)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
