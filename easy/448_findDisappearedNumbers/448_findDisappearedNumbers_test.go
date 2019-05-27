package findDisappearedNumbers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"Test 1": {
			input: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want:  []int{5, 6},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findDisappearedNumbers(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
