package mergeTrees

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

func TestMergeTrees(t *testing.T) {
	tests := map[string]struct {
		input []*TreeNode
		want  []int
	}{
		"Number 1": {
			input: []*TreeNode{
				&TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}},
				&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}},
			},
			want: []int{5, 4, 4, 3, 5, 7},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := orderByIn(mergeTrees(tc.input[0], tc.input[1]))
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
