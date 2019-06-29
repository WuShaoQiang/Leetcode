package shortestCompletingWord

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShortestCompletingWord(t *testing.T) {
	type Input struct {
		licensePlate string
		words        []string
	}
	tests := map[string]struct {
		input Input
		want  string
	}{
		"Test 1": {
			input: Input{licensePlate: "1s3 PSt", words: []string{"step", "steps", "stripe", "stepple"}},
			want:  "steps",
		},

		"Test 2": {
			input: Input{licensePlate: "1s3 456", words: []string{"looks", "pest", "stew", "show"}},
			want:  "pest",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := shortestCompletingWord(tc.input.licensePlate, tc.input.words)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
