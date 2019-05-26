package compress

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCompress(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  int
	}{
		"Test 1": {
			input: []byte("aabbbb"),
			want:  4,
		},

		"Test 2": {
			input: []byte("aabbbbbbbbbbbccc"),
			want:  7,
		},

		"Test 3": {
			input: []byte("abbbbbbbbbbbb"),
			want:  4,
		},

		"Test 4": {
			input: []byte("abbbb"),
			want:  3,
		},

		"Test 5": {
			input: []byte("aabbccc"),
			want:  6,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := compress(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
