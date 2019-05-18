package isCousins

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsCousins(t *testing.T) {
	type Input struct {
		node *TreeNode
		x    int
		y    int
	}

	tests := map[string]struct {
		input Input
		want  bool
	}{
		"Test 1": {
			input: Input{

				node: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}},
				x:    4,
				y:    5,
			},

			want: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isCousins(tc.input.node, tc.input.x, tc.input.y)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
