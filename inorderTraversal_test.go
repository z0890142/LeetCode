package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inorderTraversalTree1 = TreeNode{
	1,
	nil,
	&TreeNode{
		2,
		&TreeNode{
			3,
			nil,
			nil,
		},
		nil,
	},
}

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result []int
	}{
		{
			&inorderTraversalTree1,
			[]int{1, 3, 2},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, inorderTraversal(testItem.root))
	}

}
