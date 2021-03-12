package KthSmallestElementisaBST

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var kthSmallest1 = TreeNode{
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
	&TreeNode{
		4,
		nil,
		nil,
	},
}

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		k      int
		result int
	}{
		{
			&kthSmallest1,
			1,
			1,
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, kthSmallest(testItem.root, testItem.k))
	}

}
