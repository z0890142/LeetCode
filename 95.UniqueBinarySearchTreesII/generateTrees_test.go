package UniqueBinarySearchTreesII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []*TreeNode
var result1 = TreeNode{
	1,
	nil,
	&TreeNode{
		3,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	},
}
var result2 = TreeNode{
	3,
	&TreeNode{
		2,
		&TreeNode{
			1,
			nil,
			nil,
		},
		nil,
	},
	nil,
}
var result3 = TreeNode{
	3,
	&TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	},
	nil,
}

var result4 = TreeNode{
	2,
	&TreeNode{
		1,
		nil,
		nil,
	},
	&TreeNode{
		3,
		nil,
		nil,
	},
}
var result5 = TreeNode{
	1,
	nil,
	&TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			nil,
		},
	},
}

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		n      int
		result []*TreeNode
	}{
		{
			3,
			[]*TreeNode{&result1, &result2, &result3, &result4, &result5},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, generateTrees(testItem.n))
	}

}
