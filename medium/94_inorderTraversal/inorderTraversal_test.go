package inorderTraversal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInorderTraversal(t *testing.T) {
	tests := map[string]struct {
		input *TreeNode
		want  []int
	}{
		"Test 1": {
			input: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			want:  []int{1, 3, 2},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := inorderTraversal(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
