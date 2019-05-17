package searchBST

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func orderByIn(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	if root.Left != nil {
		res = append(res, orderByIn(root.Left)...)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		res = append(res, orderByIn(root.Right)...)
	}

	return res
}

func TestSearchBST(t *testing.T) {
	type Input struct {
		node *TreeNode
		val  int
	}
	tests := map[string]struct {
		input Input
		want  []int
	}{
		"Number 1": {
			input: Input{
				node: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}},
				val:  2,
			},
			want: []int{1, 2, 3},
		},

		"Number 2": {
			input: Input{
				node: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}},
				val:  5,
			},
			want: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := orderByIn(searchBST(tc.input.node, tc.input.val))
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
