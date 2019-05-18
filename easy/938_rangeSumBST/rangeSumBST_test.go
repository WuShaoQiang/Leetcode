package rangeSumBST

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRangeSumBST(t *testing.T) {
	type Input struct {
		node *TreeNode
		L    int
		R    int
	}

	tests := map[string]struct {
		input Input
		want  int
	}{
		"Test 1": {
			input: Input{

				node: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}},
				L:    1,
				R:    5,
			},

			want: 10,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := rangeSumBST(tc.input.node, tc.input.L, tc.input.R)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
