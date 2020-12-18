package BinaryTreeLevelOrderTraversalII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var levelOrderBottomTree1 = TreeNode{
	3,
	&TreeNode{
		9,
		nil,
		nil,
	},
	&TreeNode{
		20,
		&TreeNode{
			15,
			nil,
			nil,
		},
		&TreeNode{
			7,
			nil,
			nil,
		},
	},
}

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result [][]int
	}{
		{
			&levelOrderBottomTree1,
			[][]int{
				[]int{15, 7},
				[]int{9, 20},
				[]int{3},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, levelOrderBottom(testItem.root))
	}

}
