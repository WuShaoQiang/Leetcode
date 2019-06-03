package largeGroupPositions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLargeGroupPositions(t *testing.T) {
	tests := map[string]struct {
		input string
		want  [][]int
	}{
		"Test 1": {
			input: "abbxxxxzzy",
			want:  [][]int{[]int{3, 6}},
		},

		"Test 2": {
			input: "a",
			want:  [][]int{},
		},

		"Test 3": {
			input: "abc",
			want:  [][]int{},
		},

		"Test 4": {
			input: "abcdddeeeeaabbbcd",
			want:  [][]int{[]int{3, 5}, []int{6, 9}, []int{12, 14}},
		},

		"Test 5": {
			input: "aaa",
			want:  [][]int{[]int{0, 2}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := largeGroupPositions(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
