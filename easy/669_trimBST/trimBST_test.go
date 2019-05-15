package trimBST

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func inOrder(root *TreeNode) []int {
	res := make([]int, 0)

	if root.Left != nil {
		res = append(res, inOrder(root.Left)...)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		res = append(res, inOrder(root.Right)...)
	}

	return res
}

func TestTrimBST(t *testing.T) {
	type Input struct {
		node *TreeNode
		L    int
		R    int
	}

	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Number 2": {
			input: Input{
				node: &TreeNode{Val: 3, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4}},
				L:    1,
				R:    3,
			},
			want: []int{1, 2, 3},
		},

		"Number 1": {
			input: Input{
				node: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
				L:    1,
				R:    2,
			},
			want: []int{1, 2},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := inOrder(trimBST(tc.input.node, tc.input.L, tc.input.R))
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
