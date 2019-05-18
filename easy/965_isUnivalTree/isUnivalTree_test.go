package isUnivalTree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsUnivalTree(t *testing.T) {
	tests := map[string]struct {
		input *TreeNode
		want  bool
	}{
		"Test 1": {
			input: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},

			want: false,
		},

		"Test 2": {
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1}},

			want: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isUnivalTree(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
