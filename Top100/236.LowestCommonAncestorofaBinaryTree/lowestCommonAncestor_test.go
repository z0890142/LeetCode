package LowestCommonAncestorofaBinaryTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test1 = TreeNode{
	3,
	&test1p,
	&test1q,
}
var test1p = TreeNode{
	5,
	&TreeNode{
		6,
		nil,
		nil,
	},
	&TreeNode{
		2,
		&TreeNode{
			7,
			nil,
			nil,
		},
		&tt,
	},
}
var tt = TreeNode{
	4,
	nil,
	nil,
}
var test1q = TreeNode{
	1,
	&TreeNode{
		0,
		nil,
		nil,
	},
	&TreeNode{
		8,
		nil,
		nil,
	},
}

var test2 = TreeNode{
	3,
	&test1p,

	&TreeNode{
		1,
		&TreeNode{
			0,
			nil,
			nil,
		},
		&TreeNode{
			8,
			nil,
			nil,
		},
	},
}

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		p      *TreeNode
		q      *TreeNode
		result *TreeNode
	}{
		{
			&test1,
			&test1p,
			&test1q,
			&test1,
		},
		{
			&test2,
			&test1p,
			&tt,
			&test1p,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, lowestCommonAncestor(testItem.root, testItem.p, testItem.q))
	}

}
