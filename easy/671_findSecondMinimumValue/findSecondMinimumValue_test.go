package findSecondMinimumValue

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindSecondMinimumValue(t *testing.T) {

	tests := map[string]struct {
		input *TreeNode
		want  int
	}{
		"Number 1": {
			input: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}},

			want: 5,
		},

		"Number 2": {
			input: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}},

			want: -1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findSecondMinimumValue(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
