package averageOfLevels

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAverageOfLevels(t *testing.T) {
	tests := map[string]struct {
		input *TreeNode
		want  []float64
	}{
		"Number 1": {
			input: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want:  []float64{3, 14.5, 11},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := averageOfLevels(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}

}
