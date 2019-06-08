package sortArrayByParityII

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortArrayByParityII(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"Test 1": {
			input: []int{4, 2, 5, 7},
			want:  []int{4, 5, 2, 7},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortArrayByParityII(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
