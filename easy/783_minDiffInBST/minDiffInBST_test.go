package minDiffInBST

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinDiffInBST(t *testing.T) {

	tests := map[string]struct {
		input *TreeNode
		want  int
	}{
		"Number 1": {
			input: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},

			want: 1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := minDiffInBST(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
