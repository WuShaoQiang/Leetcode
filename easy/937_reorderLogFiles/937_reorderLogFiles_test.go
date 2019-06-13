package reorderLogFiles

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReorderLogFiles(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"Test 1": {
			input: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			want:  []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reorderLogFiles(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
