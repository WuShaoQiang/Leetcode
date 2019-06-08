package sumEvenAfterQueries

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSumEvenAfterQueries(t *testing.T) {
	type Input struct {
		A       []int
		queries [][]int
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Test 1": {
			input: Input{A: []int{1, 2, 3, 4}, queries: [][]int{[]int{1, 0}, []int{-3, 1}, []int{-4, 0}, []int{2, 3}}},
			want:  []int{8, 6, 2, 4},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sumEvenAfterQueries(tc.input.A, tc.input.queries)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
