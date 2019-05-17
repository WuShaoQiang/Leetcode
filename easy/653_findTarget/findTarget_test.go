package findTarget

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindTarget(t *testing.T) {
	type Input struct {
		node *TreeNode
		k    int
	}

	tests := map[string]struct {
		input Input
		want  bool
	}{
		"Number 1": {
			input: Input{
				node: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
				k:    9,
			},
			want: true,
		},

		"Number 2": {
			input: Input{
				node: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
				k:    28,
			},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findTarget(tc.input.node, tc.input.k)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
