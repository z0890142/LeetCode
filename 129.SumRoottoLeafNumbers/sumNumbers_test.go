package SumRoottoLeafNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		node   *TreeNode
		result int
	}{
		{
			&TreeNode{
				1,
				&TreeNode{
					2, nil, nil,
				},
				&TreeNode{
					3, nil, nil,
				},
			},
			25,
		},
		{
			&TreeNode{
				4,
				&TreeNode{
					9,
					&TreeNode{
						5,
						nil,
						nil,
					},
					&TreeNode{
						1,
						nil,
						nil,
					},
				},
				&TreeNode{
					0, nil, nil,
				},
			},
			1026,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, sumNumbers(testItem.node))

	}

}
