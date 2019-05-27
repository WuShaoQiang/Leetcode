package fib

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFib(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"Test 1": {
			input: 2,
			want:  1,
		},

		"Test 2": {
			input: 3,
			want:  2,
		},

		"Test 3": {
			input: 4,
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fib(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
