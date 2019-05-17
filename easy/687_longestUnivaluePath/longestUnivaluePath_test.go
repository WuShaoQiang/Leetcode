package longestUnivaluePath

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLongestUnivaluePath(t *testing.T) {

	tests := map[string]struct {
		input *TreeNode
		want  int
	}{
		"Number 1": {
			input: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}},

			want: 2,
		},

		"Number 2": {
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}},

			want: 2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := longestUnivaluePath(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
