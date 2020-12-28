package BinaryTreeZigzagLevelOrderTraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var zigzagLevelOrder1 = TreeNode{
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
var zigzagLevelOrder2 = TreeNode{
	1,
	&TreeNode{
		2,
		&TreeNode{
			4,
			nil,
			nil,
		},
		nil,
	},
	&TreeNode{
		3,
		nil,
		&TreeNode{
			5,
			nil,
			nil,
		},
	},
}

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result [][]int
	}{
		{
			&zigzagLevelOrder1,
			[][]int{
				[]int{3},
				[]int{20, 9},
				[]int{15, 7},
			},
		},
		{
			&zigzagLevelOrder2,
			[][]int{
				[]int{1},
				[]int{3, 2},
				[]int{4, 5},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, zigzagLevelOrder(testItem.root))
	}

}
