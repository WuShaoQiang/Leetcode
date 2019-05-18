package leafSimilar

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLeafSimilar(t *testing.T) {

	tests := map[string]struct {
		input []*TreeNode
		want  bool
	}{
		"Test 1": {
			input: []*TreeNode{
				&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},
				&TreeNode{Val: 3, Left: &TreeNode{Val: 10, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},
			},
			want: true,
		},

		"Test 2": {
			input: []*TreeNode{
				&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},
				&TreeNode{Val: 3, Left: &TreeNode{Val: 10, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},
			},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := leafSimilar(tc.input[0], tc.input[1])
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
