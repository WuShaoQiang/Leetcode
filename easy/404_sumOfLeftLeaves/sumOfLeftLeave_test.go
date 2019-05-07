package sumOfLeftLeaves

import (
	"fmt"
	"testing"
)

var (
	input = &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
)

func TestSumOfLeftLeaves(t *testing.T) {
	fmt.Println(sumOfLeftLeaves(input))
}
