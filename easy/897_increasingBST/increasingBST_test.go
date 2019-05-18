package increasingBST

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

func TestIncreasingBST(t *testing.T) {
	tests := map[string]struct {
		input *TreeNode
		want  []int
	}{
		"Number 1": {
			input: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},

			want: []int{1, 2, 3, 4},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := orderByIn(increasingBST(tc.input))
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
