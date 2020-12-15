package PathSumII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pathSumResult2 = TreeNode{
	5,
	&TreeNode{
		4,
		&TreeNode{
			11,
			&TreeNode{
				7,
				nil,
				nil,
			},
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		nil,
	},
	&TreeNode{
		8,
		&TreeNode{
			13,
			nil,
			nil,
		},
		&TreeNode{
			4,
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
	},
}

func Test_pathSum(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		sum    int
		result [][]int
	}{
		{
			&pathSumResult2,
			22,
			[][]int{
				[]int{5, 4, 11, 2},
				[]int{5, 8, 4, 5},
			},
		},
	}
	for _, testItem := range tests {

		assert.Equal(t, testItem.result, pathSum(testItem.root, testItem.sum))

	}

}
