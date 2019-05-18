package sumRootToLeaf

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsUnivalTree(t *testing.T) {
	tests := map[string]struct {
		input *TreeNode
		want  int
	}{
		"Test 1": {
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}},

			want: 22,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sumRootToLeaf(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
