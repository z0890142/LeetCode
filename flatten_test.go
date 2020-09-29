package main

import (
	"testing"
)

var flattenResult = TreeNode{
	1,
	&TreeNode{
		2,
		&TreeNode{
			3,
			nil,
			nil,
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	},
	&TreeNode{
		5,
		nil,
		&TreeNode{
			6,
			nil,
			nil,
		},
	},
}

func Test_flatten(t *testing.T) {
	tests := []struct {
		root *TreeNode
	}{
		{
			&flattenResult,
		},
	}
	for _, testItem := range tests {
		flatten(testItem.root)

	}

}
